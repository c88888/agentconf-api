package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/n9e"
)

type SearchPluginTplParms struct {
	model.PluginTpl
	request.PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}

type PluginTplIdent struct {
	Ident string `json:"ident" xml:"ident" form:"ident"`
}
