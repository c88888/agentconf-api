package service

type ServiceGroup struct {
	BusiGroupService
	CollectRuleService
	ConfigService
	TargetService
	PluginTplService
}

var ServiceGroupApp = new(ServiceGroup)
