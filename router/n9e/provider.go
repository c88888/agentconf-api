package n9e

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProviderRouter struct{}

func (p *ProviderRouter) InitProviderRouter(Router *gin.RouterGroup) {
	providerRouter := Router.Group("config").Use(middleware.SignAuth())
	providerApi := v1.ApiGroupApp.N9eApiGroup.ConfigApi
	{
		providerRouter.GET("categraf", providerApi.GetCategrafConfigByHost) // 获取采集配置
	}
}
