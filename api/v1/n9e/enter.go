package n9e

import "github.com/flipped-aurora/gin-vue-admin/server/service/n9e"

type ApiGroup struct {
	BusiGroupApi
	CollectRuleApi
	ConfigApi
	TargetApi
	PluginTplApi
}

var (
	busigroupService   = service.ServiceGroupApp.BusiGroupService
	collectRuleService = service.ServiceGroupApp.CollectRuleService
	configService      = service.ServiceGroupApp.ConfigService
	targetService      = service.ServiceGroupApp.TargetService
	pluginTplService   = service.ServiceGroupApp.PluginTplService
)
