package model

type BusiGroup struct {
	ID          int64  `json:"id" gorm:"id"`
	Name        string `json:"name" gorm:"name"`
	LabelEnable int8   `json:"label_enable" gorm:"label_enable"`
	LabelValue  string `json:"label_value" gorm:"label_value"` // if label_enable: label_value can not be blank
	CreateAt    int64  `json:"create_at" gorm:"create_at"`
	CreateBy    string `json:"create_by" gorm:"create_by"`
	UpdateAt    int64  `json:"update_at" gorm:"update_at"`
	UpdateBy    string `json:"update_by" gorm:"update_by"`

	UserGroups []UserGroupWithPermFlag `json:"user_groups" gorm:"-"`
}
type UserGroupWithPermFlag struct {
	UserGroup *UserGroup `json:"user_group"`
	PermFlag  string     `json:"perm_flag"`
}

func (*BusiGroup) TableName() string {
	return "busi_group"
}
