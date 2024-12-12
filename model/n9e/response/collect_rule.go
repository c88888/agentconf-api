package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/n9e"
)

type CollectRuleResponse struct {
	CollectRule model.CollectRule `json:"collect_rule"`
}

type CollectRuleProviderResponse struct {
	// version is signature/md5 of current Config, server side should deal with the Version calculate
	Version string `json:"version"`

	// ConfigMap (InputName -> Config), if version is identical, server side can set Config to nil
	Configs map[string]map[string]*ConfigWithFormat `json:"configs"`
}

type ConfigFormat string

type ConfigWithFormat struct {
	Config string       `json:"config"`
	Format ConfigFormat `json:"format"`
	//checkSum string       `json:"-"` //客户端使用
}
