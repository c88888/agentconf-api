package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/n9e"
)

type SearchTargetParams struct {
	model.Target
	request.PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}

type TargetParams struct {
	ID       int64  `json:"id" gorm:"id" binding:"required"`
	GroupId  int64  `json:"group_id" gorm:"group_id" binding:"gte=0,lte=2000"` // busi group id
	Ident    string `json:"ident" gorm:"ident"`                                // target id
	Note     string `json:"note" gorm:"note"`                                  // append to alert event as field
	Tags     string `json:"-" gorm:"-"`                                        // append to series data as tags, split by space, append external space at suffix
	UpdateAt int64  `json:"update_at" gorm:"update_at"`

	GroupObj *model.BusiGroup  `json:"group_obj" gorm:"-"`
	Cluster  string            `json:"cluster"`
	TagsJSON []string          `json:"tags" gorm:"-"`
	TagsMap  map[string]string `json:"-" gorm:"-"` // internal use, append tags to series

	TargetUp    float64  `json:"target_up" gorm:"-"`
	LoadPerCore float64  `json:"load_per_core" gorm:"-"`
	MemUtil     float64  `json:"mem_util" gorm:"-"`
	DiskUtil    float64  `json:"disk_util" gorm:"-"`
	Ids         []string `json:"ids" gorm:"-"`
}

type TargetIdsParams struct {
	GroupId int64    `json:"group_id" gorm:"group_id" binding:"gte=0,lte=2000"`
	Ids     []string `json:"ids" gorm:"ids" binding:"required"`
}

type TargetTagsParams struct {
	GroupId int64    `json:"group_id" gorm:"group_id" binding:"gte=0,lte=2000"`
	Tags    []string `json:"tags" gorm:"tags" binding:"required"`
}
