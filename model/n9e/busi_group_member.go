package model

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type BusiGroupMember struct {
	BusiGroupId int64  `json:"busi_group_id"`
	UserGroupId int64  `json:"user_group_id"`
	PermFlag    string `json:"perm_flag"`
}

func (BusiGroupMember) TableName() string {
	return "busi_group_member"
}

func BusiGroupIds(userGroupIds []int64, permFlag ...string) ([]int64, error) {
	if len(userGroupIds) == 0 {
		return []int64{}, nil
	}

	session := global.MustGetGlobalDBByDBName("n9e").Model(&BusiGroupMember{}).Where("user_group_id in ?", userGroupIds)
	if len(permFlag) > 0 {
		session = session.Where("perm_flag=?", permFlag[0])
	}

	var ids []int64
	err := session.Pluck("busi_group_id", &ids).Error
	return ids, err
}
