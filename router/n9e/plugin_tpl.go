package n9e

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PluginTplRouter struct{}

func (*PluginTplRouter) InitPluginTplRouter(Router *gin.RouterGroup) {
	pluginTplRouter := Router.Group("plugin").Use(middleware.OperationRecord())
	pluginTplApi := v1.ApiGroupApp.N9eApiGroup.PluginTplApi
	{
		pluginTplRouter.POST("tpl", pluginTplApi.CreatePluginTpl)   // 创建插件模板
		pluginTplRouter.PUT("tpl", pluginTplApi.UpdatePluginTpl)    // 更新插件模板
		pluginTplRouter.DELETE("tpl", pluginTplApi.DeletePluginTpl) // 删除插件模板
	}
	{
		pluginTplRouter.GET("list", pluginTplApi.PluginTplList)             //获取列表
		pluginTplRouter.GET("getById", pluginTplApi.GetPluginTplById)       //根据id获取单一插件模板
		pluginTplRouter.GET("getByIdent", pluginTplApi.GetPluginTplByIdent) //根据ident标识获取单一插件模板
	}

}
