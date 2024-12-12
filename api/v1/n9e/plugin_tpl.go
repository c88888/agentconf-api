package n9e

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/n9e"
	pluginTplReq "github.com/flipped-aurora/gin-vue-admin/server/model/n9e/request"
	pluginTplRes "github.com/flipped-aurora/gin-vue-admin/server/model/n9e/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"time"
)

type PluginTplApi struct{}

// CreatePluginTpl
// @Tags PluginTplApi
// @Summary 创建插件模板
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      model.PluginTpl
// @Success   200   {object}  response.Response{msg=string}  "创建插件"
// @Router    /plugin/tpl [post]
func (*PluginTplApi) CreatePluginTpl(c *gin.Context) {
	var pt model.PluginTpl
	if err := c.ShouldBindJSON(&pt); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if pt.Ident == "" {
		response.FailWithMessage("插件标识错误", c)
	}

	if err := utils.Verify(pt, utils.PluginTplVerify); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	pt.CreatedBy = utils.GetUserInfo(c).NickName
	pt.UpdatedBy = utils.GetUserInfo(c).NickName
	now := time.Now().Unix()
	pt.CreatedAt = now
	pt.UpdatedAt = now
	err := pluginTplService.CreatePluginTpl(pt)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// UpdatePluginTpl
// @Tags PluginTplApi
// @Summary 创建插件模板
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      model.PluginTpl
// @Success   200   {object}  response.Response{msg=string}  "更新插件模板"
// @Router    /plugin/tpl [put]
func (b *PluginTplApi) UpdatePluginTpl(c *gin.Context) {
	var pt model.PluginTpl
	err := c.ShouldBindJSON(&pt)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err = utils.Verify(pt, utils.PluginTplIdVerify); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	pt.UpdatedBy = utils.GetUserInfo(c).NickName
	now := time.Now().Unix()
	pt.UpdatedAt = now
	err = pluginTplService.UpdatePluginTpl(pt)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeletePluginTpl
// @Tags PluginTplApi
// @Summary 删除插件模板
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      model.PluginTpl
// @Success   200   {object}  response.Response{msg=string}  "删除插件模板"
// @Router    /plugin/tpl [delete]
func (*PluginTplApi) DeletePluginTpl(c *gin.Context) {
	var pt model.PluginTpl
	if err := c.ShouldBindJSON(&pt); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := utils.Verify(pt, utils.PluginTplIdVerify); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := pluginTplService.DeletePluginTpl(pt)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// PluginTplList
// @Tags PluginTplApi
// @Summary 获取插件模板列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      model.PluginTpl
// @Param data body request.PageInfo true "分页获取API列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取API列表,返回包括列表,总数,页码,每页数量"
// @Router  /plugin/list [get]
func (b *PluginTplApi) PluginTplList(c *gin.Context) {
	var pageInfo pluginTplReq.SearchPluginTplParms
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	ptList, total, err := pluginTplService.GetPluginTplList(pageInfo.PluginTpl, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     ptList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags PluginTplApi
// @Summary 根据id获取PluginTpl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {object} response.Response{data=systemRes.SysAPIResponse} "根据id获取api,返回包括api详情"
// @Router  /plugin/getById [get]
func (b *PluginTplApi) GetPluginTplById(c *gin.Context) {
	var idInfo request.GetById
	if err := c.ShouldBindQuery(&idInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	pt, err := pluginTplService.GetPluginTplById(int64(idInfo.ID))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pluginTplRes.PluginTplResponse{PluginTpl: pt}, "获取成功", c)
	}
}

// @Tags PluginTplApi
// @Summary 根据ident标识获取PluginTpl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ident true "根据ident获取PluginTpl"
// @Success 200 {object} response.Response{data=pluginTplRes.PluginTplResponse} "根据ident获取PluginTpl,返回包括PluginTpl详情"
// @Router  /plugin/getById [get]
func (b *PluginTplApi) GetPluginTplByIdent(c *gin.Context) {
	var pt pluginTplReq.PluginTplIdent
	err := c.ShouldBindQuery(&pt)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(pt, utils.PluginTplIdentVerify)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	p, err := pluginTplService.GetPluginTplByIdent(pt.Ident)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(p, "获取成功", c)
	}
}
