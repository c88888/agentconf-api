package n9e

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/n9e"
	targetReq "github.com/flipped-aurora/gin-vue-admin/server/model/n9e/request"
	n9eRes "github.com/flipped-aurora/gin-vue-admin/server/model/n9e/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type TargetApi struct{}

// 修改or移出业务组
func (*TargetApi) MoveTargetFromBusigroup(c *gin.Context) {
	var targetParms targetReq.TargetParams
	if err := c.ShouldBindJSON(&targetParms); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var targetInfo model.Target
	targetInfo.GroupId = targetParms.GroupId
	targetInfo.ID = targetParms.ID
	if err := targetService.MoveTargetToBusigroup(targetInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func (*TargetApi) GetTargetsByGroupIDAndIdents(c *gin.Context) {
	var targetParams targetReq.TargetIdsParams
	if err := c.ShouldBindJSON(&targetParams); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	ids := targetParams.Ids
	gid := targetParams.GroupId
	targetList, err := targetService.GetTargetsByGroupIDAndIdents(gid, ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(targetList, c)
}

func (*TargetApi) GetTargetsByGroupIDAndTags(c *gin.Context) {
	var targetParams targetReq.TargetTagsParams
	if err := c.ShouldBindJSON(&targetParams); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	tags := targetParams.Tags
	gid := targetParams.GroupId
	targetList, err := targetService.GetTargetsByGroupIDAndTags(gid, tags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(targetList, c)
}

func (*TargetApi) GetTargetList(c *gin.Context) {
	mins := 2
	start := time.Now().Add(-time.Hour).Format(time.RFC3339)
	end := time.Now().Format(time.RFC3339)
	step := "1m"
	var pageInfo targetReq.SearchTargetParams
	if err := c.ShouldBindJSON(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := targetService.GetTargetList(&pageInfo.Target, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)

	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	}

	if err == nil {
		cache := make(map[int64]*model.BusiGroup)
		targetsMap := make(map[string]*model.Target)
		for i := 0; i < len(list); i++ {
			global.GVA_LOG.Sugar().Debugln(targetService.FillGroup(cache, list[i]))
			targetsMap[list[i].Cluster+list[i].Ident] = list[i]
		}

		// query LoadPerCore / MemUtil / TargetUp / DiskUsedPercent from prometheus
		// map key: cluster, map value: ident list
		targets := make(map[string][]string)
		for i := 0; i < len(list); i++ {
			targets[list[i].Cluster] = append(targets[list[i].Cluster], list[i].Ident)
		}

		for cluster := range targets {
			targetArr := targets[cluster]
			if len(targetArr) == 0 {
				continue
			}

			targetRe := strings.Join(targetArr, "|")
			valuesMap := make(map[string]map[string]float64)

			for metric, ql := range utils.StructToMap(global.GVA_CONFIG.Target) {
				promql := fmt.Sprintf(ql.(string), targetRe, mins)
				values, err := utils.InstantQuery("query", promql, start, end, step)
				if err != nil {
					global.GVA_LOG.Sugar().Warnln(err)
				}
				valuesMap[metric] = values
			}
			// handle values
			for metric, values := range valuesMap {
				for ident := range values {
					mapkey := cluster + ident
					if t, ok := targetsMap[mapkey]; ok {
						switch metric {
						case "LoadPerCore":
							t.LoadPerCore = values[ident]
						case "MemUtil":
							t.MemUtil = values[ident]
						case "TargetUp":
							t.TargetUp = values[ident]
						case "DiskUtil":
							t.DiskUtil = values[ident]
						}
					}
				}
			}
		}
	}
	c.JSON(http.StatusOK, n9eRes.TargetListResponse{Code: 0, List: list, Total: total})
}

/*
req:
{
   "group_id":25,
   "page":1,
   "pageSize":10,
   "orderKey":"id"
}

response:
{
    "list": [
        {
            "id": 46,
            "group_id": 25,
            "ident": "zy-wy-dev-yiqiyoo-db-01",
            "note": "",
            "update_at": 1658941128,
            "group_obj": {
                "id": 25,
                "name": "yiqiyoo-saas",
                "label_enable": 1,
                "label_value": "yiqiyoo-saas",
                "create_at": 1658941127,
                "create_by": "822032277",
                "update_at": 1658941127,
                "update_by": "822032277"
            },
            "cluster": "Default",
            "tags": null,
            "target_up": 1,
            "load_per_core": 0.1025,
            "mem_util": 84.149028867655,
            "disk_util": 33.9875436409
        },
        {
            "id": 47,
            "group_id": 25,
            "ident": "zy-js-pro-yiqiyoo-db-01",
            "note": "",
            "update_at": 1684720205,
            "group_obj": {
                "id": 25,
                "name": "yiqiyoo-saas",
                "label_enable": 1,
                "label_value": "yiqiyoo-saas",
                "create_at": 1658941127,
                "create_by": "822032277",
                "update_at": 1658941127,
                "update_by": "822032277"
            },
            "cluster": "Default",
            "tags": null,
            "target_up": 1,
            "load_per_core": 0.135,
            "mem_util": 57.1035433573,
            "disk_util": 29.886221377242
        },
        {
            "id": 341,
            "group_id": 25,
            "ident": "zy-wh-dev-yiqiyoo-game-01",
            "note": "",
            "update_at": 1689662867,
            "group_obj": {
                "id": 25,
                "name": "yiqiyoo-saas",
                "label_enable": 1,
                "label_value": "yiqiyoo-saas",
                "create_at": 1658941127,
                "create_by": "822032277",
                "update_at": 1658941127,
                "update_by": "822032277"
            },
            "cluster": "Default",
            "tags": null,
            "target_up": 1,
            "load_per_core": 0.0425,
            "mem_util": 19.329954319400002,
            "disk_util": 8.37827053456
        }
    ],
    "total": 3
}
*/
