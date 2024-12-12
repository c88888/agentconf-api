package n9e

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BusiGroupRouter struct{}

func (u *BusiGroupRouter) InitBusiGroupRouter(Router *gin.RouterGroup) {
	busigroupRouter := Router.Group("busigroup").Use(middleware.OperationRecord())
	busigroupApi := v1.ApiGroupApp.N9eApiGroup.BusiGroupApi
	{
		busigroupRouter.GET("getById", busigroupApi.GetBusiGroupById) // 单个获取
		busigroupRouter.POST("list", busigroupApi.BusiGroupList)      // 分页获取列表
	}
}
