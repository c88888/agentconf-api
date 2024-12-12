package model

type UserGroup struct {
	Id       int64   `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name"`
	Note     string  `json:"note"`
	CreateAt int64   `json:"create_at"`
	CreateBy string  `json:"create_by"`
	UpdateAt int64   `json:"update_at"`
	UpdateBy string  `json:"update_by"`
	UserIds  []int64 `json:"-" gorm:"-"`
}

func (ug *UserGroup) TableName() string {
	return "user_group"
}
