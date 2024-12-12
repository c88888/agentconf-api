package response

import model "github.com/flipped-aurora/gin-vue-admin/server/model/n9e"

type PluginTplResponse struct {
	*model.PluginTpl `json:"plugin_tpl"`
}
