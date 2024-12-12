package n9e

import (
	"github.com/gin-gonic/gin"
	"net/http"

	n9eRes "github.com/flipped-aurora/gin-vue-admin/server/model/n9e/response"
)

type ConfigApi struct{}

// @Tags ConfigApi
// @Summary agent获取采集配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {object} n9eRes.CollectRuleProviderResponse "根据id获取api,返回包括api详情"
// @Router  /config/categraf [get]
func (ca *ConfigApi) GetCategrafConfigByHost(c *gin.Context) {
	version := c.Query("version")
	agentHostname := c.Query("agent_hostname")
	data, err := configService.CollectRuleProvide(version, agentHostname)
	if data != nil && err == nil {
		c.JSON(http.StatusOK, data)
		return
	}

	c.JSON(http.StatusOK, n9eRes.CollectRuleProviderResponse{Version: "", Configs: nil})
}
