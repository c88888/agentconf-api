package model

type Target struct {
	ID       int64  `json:"id" gorm:"id"`
	GroupId  int64  `json:"group_id" gorm:"group_id"` // busi group id
	Ident    string `json:"ident" gorm:"ident"`       // target id
	Note     string `json:"note" gorm:"note"`         // append to alert event as field
	Tags     string `json:"tags" gorm:"tags"`         // append to series data as tags, split by space, append external space at suffix
	UpdateAt int64  `json:"update_at" gorm:"update_at"`

	GroupObj *BusiGroup `json:"group_obj" gorm:"-"`
	Cluster  string     `json:"cluster"`
	//TagsJSON []string          `json:"tags" gorm:"-"`
	TagsMap map[string]string `json:"-" gorm:"-"` // internal use, append tags to series

	TargetUp    float64 `json:"target_up" gorm:"-"`
	LoadPerCore float64 `json:"load_per_core" gorm:"-"`
	MemUtil     float64 `json:"mem_util" gorm:"-"`
	DiskUtil    float64 `json:"disk_util" gorm:"-"`
}

func (*Target) TableName() string {
	return "target"
}
