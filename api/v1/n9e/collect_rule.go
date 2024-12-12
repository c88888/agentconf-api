package n9e

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/n9e"
	n9eReq "github.com/flipped-aurora/gin-vue-admin/server/model/n9e/request"
	n9eRes "github.com/flipped-aurora/gin-vue-admin/server/model/n9e/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type CollectRuleApi struct{}

// @Tags CollectRuleApi
// @Summary 分页获取BusiGroupApi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取API列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取API列表,返回包括列表,总数,页码,每页数量"
// @Router  /collect/list [post]
func (cra *CollectRuleApi) CollectRuleList(c *gin.Context) {
	var pageInfo n9eReq.SearchCollectRuleParms
	if err := c.ShouldBindJSON(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if pageInfo.GroupId == 0 {
		response.FailWithMessage("业务组id错误", c)
	}
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	customerList, total, err := collectRuleService.GetCollectRuleList(pageInfo.CollectRule, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
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
}

// @Tags CollectRuleApi
// @Summary 根据id获取CollectRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {object} response.Response{data=systemRes.SysAPIResponse} "根据id获取api,返回包括api详情"
// @Router  /collect/getById [post]
func (cra *CollectRuleApi) GetCollectRuleById(c *gin.Context) {
	var idInfo request.GetById
	if err := c.ShouldBindQuery(&idInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(idInfo, utils.PluginTplIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cr, err := collectRuleService.GetCollectRuleById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(n9eRes.CollectRuleResponse{CollectRule: cr}, "获取成功", c)
	}
}

/*
{"group_name":"ai-nebula","group_id":70,"name":"sss"}
*/
func (cra *CollectRuleApi) CreateCollectRule(c *gin.Context) {
	var cr model.CollectRule
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if cr.GroupId == 0 {
		response.FailWithMessage("业务组id错误", c)
	}
	if err := utils.Verify(cr, utils.CollectRuleVerify); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	cr.CreateBy = utils.GetUserInfo(c).NickName
	cr.UpdateBy = utils.GetUserInfo(c).NickName
	now := time.Now().Unix()
	cr.CreateAt = now
	cr.UpdateAt = now
	err := collectRuleService.CreateCollectRule(cr)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(n9eRes.CollectRuleResponse{CollectRule: cr}, "创建成功", c)
	}
}

func (cra *CollectRuleApi) GetCollectRuleByName(c *gin.Context) {
	var idInfo request.GetById
	if err := c.ShouldBindQuery(&idInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(idInfo, utils.PluginTplIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cr, err := collectRuleService.GetCollectRuleById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(n9eRes.CollectRuleResponse{CollectRule: cr}, "获取成功", c)
	}
}

func (cra *CollectRuleApi) UpdateCollectRule(c *gin.Context) {
	var cr model.CollectRule
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if cr.GroupId == 0 {
		response.FailWithMessage("业务组id错误", c)
	}
	if err := utils.Verify(cr, utils.CollectRuleVerify); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	cr.UpdateBy = utils.GetUserInfo(c).NickName
	now := time.Now().Unix()
	cr.UpdateAt = now
	err := collectRuleService.UpdateCollectRule(cr)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(n9eRes.CollectRuleResponse{CollectRule: cr}, "更新成功", c)
	}
}

func (cra *CollectRuleApi) DeleteCollectRule(c *gin.Context) {
	var cr model.CollectRule
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if cr.GroupId == 0 {
		response.FailWithMessage("业务组id错误", c)
	}
	if err := utils.Verify(cr, utils.IdVerify); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := collectRuleService.DeleteCollectRule(cr)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithDetailed(n9eRes.CollectRuleResponse{CollectRule: cr}, "删除成功", c)
	}
}
