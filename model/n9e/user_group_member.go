package model

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type UserGroupMember struct {
	GroupId int64
	UserId  int64
}

func (UserGroupMember) TableName() string {
	return "user_group_member"
}

func MyGroupIds(userId int64) ([]int64, error) {
	var ids []int64
	err := global.MustGetGlobalDBByDBName("n9e").Model(&UserGroupMember{}).Where("user_id=?", userId).Pluck("group_id", &ids).Error
	return ids, err
}

func MemberIds(groupId int64) ([]int64, error) {
	var ids []int64
	err := global.MustGetGlobalDBByDBName("n9e").Model(&UserGroupMember{}).Where("group_id=?", groupId).Pluck("user_id", &ids).Error
	return ids, err
}

func UserGroupMemberCount(where string, args ...interface{}) (int64, error) {
	var total int64
	db := global.MustGetGlobalDBByDBName("n9e").Model(&UserGroupMember{}).Where(where, args...)
	err := db.Count(&total).Error
	return total, err

}

func UserGroupMemberAdd(groupId, userId int64) error {
	num, err := UserGroupMemberCount("user_id=? and group_id=?", userId, groupId)
	if err != nil {
		return err
	}

	if num > 0 {
		// already exists
		return nil
	}

	obj := UserGroupMember{
		GroupId: groupId,
		UserId:  userId,
	}

	return global.MustGetGlobalDBByDBName("n9e").Model(&UserGroupMember{}).Create(obj).Error
}

func UserGroupMemberDel(groupId int64, userIds []int64) error {
	if len(userIds) == 0 {
		return nil
	}

	return global.MustGetGlobalDBByDBName("n9e").Where("group_id = ? and user_id in ?", groupId, userIds).Delete(&UserGroupMember{}).Error
}

func UserGroupMemberGetAll() ([]UserGroupMember, error) {
	var lst []UserGroupMember
	err := global.MustGetGlobalDBByDBName("n9e").Find(&lst).Error
	return lst, err
}
