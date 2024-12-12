package n9e

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/n9e"
	busigroupReq "github.com/flipped-aurora/gin-vue-admin/server/model/n9e/request"
	busigroupRes "github.com/flipped-aurora/gin-vue-admin/server/model/n9e/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

type BusiGroupApi struct{}

// @Tags BusiGroupApi
// @Summary 分页获取BusiGroupApi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取API列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取API列表,返回包括列表,总数,页码,每页数量"
// @Router  /busigroup/list [post]
func (b *BusiGroupApi) BusiGroupList(c *gin.Context) {
	var pageInfo busigroupReq.SearchBusiGroupParms
	if err := c.ShouldBindJSON(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// admin or 业务组成员
	var u *model.User
	uname := utils.GetUserInfo(c).Username
	err := global.MustGetGlobalDBByDBName("n9e").Where("username = ?", uname).First(&u).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	for _, v := range global.GVA_CONFIG.N9eAdmin {
		if u.Username == v {
			u.RolesLst = strings.Split(u.Roles, " ")
		}
	}

	lst, total, err := u.BusiGroups(pageInfo.PageSize, "", false)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     lst,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
	/*
		customerList, total, err := busigroupService.GetBusiGroupList(pageInfo.BusiGroup, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
		if err != nil {
			global.GVA_LOG.Error("获取失败!", zap.Error(err))
			response.FailWithMessage("获取失败"+err.Error(), c)
		} else {
			response.OkWithDetailed(response.PageResult{
				List:     customerList,
				Total:    total,
				Page:     pageInfo.Page,
				PageSize: pageInfo.PageSize,
			}, "获取成功", c)
		}
	*/
}

// @Tags BusiGroupApi
// @Summary 根据id获取BusiGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {object} busigroupRes.BusiGroupResponse{Busigroup=*model.BusiGroup} "根据id获取api,返回包括api详情"
// @Router  /busigroup/getById [post]
func (b *BusiGroupApi) GetBusiGroupById(c *gin.Context) {
	var idInfo request.GetById
	if err := c.ShouldBindQuery(&idInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	bg, err := busigroupService.GetBusiGroupById(int64(idInfo.ID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(busigroupRes.BusiGroupResponse{Busigroup: bg}, "获取成功", c)
	}
}
