package n9e

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TargetRouter struct{}

func (u *TargetRouter) InitTargetRouter(Router *gin.RouterGroup) {
	targetRouter := Router.Group("target").Use(middleware.OperationRecord())
	targetApi := v1.ApiGroupApp.N9eApiGroup.TargetApi
	{
		targetRouter.POST("list", targetApi.GetTargetList)           // 分页获取列表
		targetRouter.POST("move", targetApi.MoveTargetFromBusigroup) //移出业务组
		targetRouter.POST("getByTags", targetApi.GetTargetsByGroupIDAndTags)
		targetRouter.POST("getByIdents", targetApi.GetTargetsByGroupIDAndIdents)
	}
}
