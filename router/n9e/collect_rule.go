package n9e

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CollectRuleRouter struct{}

func (*CollectRuleRouter) InitCollectRuleRouter(Router *gin.RouterGroup) {
	collectRuleRouter := Router.Group("rule").Use(middleware.OperationRecord())
	collectRuleApi := v1.ApiGroupApp.N9eApiGroup.CollectRuleApi
	{
		collectRuleRouter.POST("list", collectRuleApi.CollectRuleList)      // 获取采集配置列表
		collectRuleRouter.GET("getById", collectRuleApi.GetCollectRuleById) //根据id获取采集配置
		collectRuleRouter.POST("collect", collectRuleApi.CreateCollectRule)
		collectRuleRouter.PUT("collect", collectRuleApi.UpdateCollectRule)
		collectRuleRouter.DELETE("collect", collectRuleApi.DeleteCollectRule)
	}

}
