package model

import (
	"github.com/didi/nightingale/v5/src/pkg/ormx"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/pkg/errors"
)

const (
	AdminRole = "Admin"
)

type User struct {
	Id         int64        `json:"id" gorm:"primaryKey"`
	Username   string       `json:"username"`
	Nickname   string       `json:"nickname"`
	Password   string       `json:"-"`
	Phone      string       `json:"phone"`
	Email      string       `json:"email"`
	Portrait   string       `json:"portrait"`
	Roles      string       `json:"-"`              // 这个字段写入数据库
	RolesLst   []string     `json:"roles" gorm:"-"` // 这个字段和前端交互
	Contacts   ormx.JSONObj `json:"contacts"`       // 内容为 map[string]string 结构
	Maintainer int          `json:"maintainer"`     // 是否给管理员发消息 0:not send 1:send
	CreateAt   int64        `json:"create_at"`
	CreateBy   string       `json:"create_by"`
	UpdateAt   int64        `json:"update_at"`
	UpdateBy   string       `json:"update_by"`
	Admin      bool         `json:"admin" gorm:"-"` // 方便前端使用
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) IsAdmin() bool {
	for i := 0; i < len(u.RolesLst); i++ {
		if u.RolesLst[i] == AdminRole {
			return true
		}
	}
	return false
}

func (u *User) BusiGroups(limit int, query string, all ...bool) ([]BusiGroup, int64, error) {
	session := global.MustGetGlobalDBByDBName("n9e").Order("name").Limit(limit)
	var lst []BusiGroup
	var total int64
	if u.IsAdmin() || (len(all) > 0 && all[0]) {
		db := session.Where("name like ?", "%"+query+"%").Find(&lst)
		err := db.Count(&total).Error
		if err != nil {
			return lst, total, err
		}
		return lst, total, err
	}

	userGroupIds, err := MyGroupIds(u.Id)
	if err != nil {
		return nil, 0, errors.WithMessage(err, "failed to get MyGroupIds")
	}

	busiGroupIds, err := BusiGroupIds(userGroupIds)
	if err != nil {
		return nil, 0, errors.WithMessage(err, "failed to get BusiGroupIds")
	}

	if len(busiGroupIds) == 0 {
		return lst, 0, nil
	}

	db := session.Where("id in ?", busiGroupIds).Where("name like ?", "%"+query+"%").Find(&lst)
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return lst, total, err
}
