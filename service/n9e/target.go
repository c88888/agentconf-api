package service

import (
	"fmt"
	"strings"
	"sync"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/n9e"
)

type TargetService struct{}

func (*TargetService) DeleteTarget(t model.Target) error {
	return global.MustGetGlobalDBByDBName("n9e").Delete(&t).Error
}

func (*TargetService) UpdateTarget(t model.Target) error {
	return global.MustGetGlobalDBByDBName("n9e").
		Where("id = ?", t.ID).
		First(&model.Target{}).
		Updates(&t).Error
}

func (*TargetService) MoveTargetToBusigroup(t model.Target) error {
	return global.MustGetGlobalDBByDBName("n9e").
		Model(&model.Target{}).
		Where("id = ?", t.ID).
		Update("group_id", t.GroupId).Error
}

// @function: GetTargetList
// @description: 分页获取列表
// @param: info request.PageInfo
// @return: list []*model.Target, total int64, err error
func (t *TargetService) GetTargetList(tg *model.Target, info request.PageInfo, order string, desc bool,
) (list []*model.Target, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MustGetGlobalDBByDBName("n9e").Model(&model.Target{})
	var targetLists []*model.Target

	db = db.Where("group_id=?", tg.GroupId)

	err = db.Count(&total).Error
	if err != nil {
		return targetLists, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			orderMap := make(map[string]bool)
			orderMap["id"] = true
			orderMap["name"] = true
			if orderMap[order] {
				if desc {
					OrderStr = order + " desc"
				} else {
					OrderStr = order
				}
			} else {
				err = fmt.Errorf("非法的排序字段: %v", order)
				return targetLists, total, err
			}
			err = db.Order(OrderStr).Find(&targetLists).Error
		} else {
			err = db.Order("id").Find(&targetLists).Error
		}
	}

	// 用带索引的map来确保排序顺序
	resultMap := make(map[int]*model.Target)
	var mu sync.Mutex
	var wg sync.WaitGroup

	if len(targetLists) >= 1 {
		for index, v := range targetLists {
			wg.Add(1)
			go func(index int, v *model.Target) {
				defer wg.Done()
				mu.Lock()
				resultMap[index] = v
				mu.Unlock()
			}(index, v)
		}
	}

	wg.Wait()

	data := make([]*model.Target, len(resultMap))
	for i := 0; i < len(resultMap); i++ {
		data[i] = resultMap[i]
	}
	return data, total, err
}

func (t *TargetService) FillGroup(cache map[int64]*model.BusiGroup, mt *model.Target) error {
	//if mt.GroupId <= 0 {
	//	return nil
	//}

	bg, has := cache[mt.GroupId]
	if has {
		mt.GroupObj = bg
		return nil
	}

	bg, err := ServiceGroupApp.BusiGroupService.GetBusiGroupById(mt.GroupId)
	if err != nil {
		return err
	}

	mt.GroupObj = bg
	cache[mt.GroupId] = bg
	return nil
}

func (*TargetService) GetTargetsByGroupIDAndTags(groupID int64, tags []string) ([]model.Target, error) {
	var targets []model.Target
	groupQuery := global.MustGetGlobalDBByDBName("n9e").Model(&model.Target{}).Where("group_id = ?", groupID)
	if err := groupQuery.Find(&targets).Error; err != nil {
		return nil, err
	}

	if len(tags) == 0 {
		return targets, nil
	}
	var filteredTargets []model.Target

	for _, t := range targets {
		matches := true
		for _, tag := range tags {
			if !strings.Contains(t.Tags, tag) {
				matches = false
				break
			}
		}
		if matches {
			filteredTargets = append(filteredTargets, t)
		}
	}
	return filteredTargets, nil
}

func (*TargetService) GetTargetsByGroupIDAndIdents(groupID int64, ids []string) ([]model.Target, error) {
	var targets []model.Target
	groupQuery := global.MustGetGlobalDBByDBName("n9e").
		Model(&model.Target{}).Where("group_id = ?", groupID)

	if err := groupQuery.Find(&targets).Error; err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return targets, nil
	}
	var filteredTargets []model.Target

	for _, target := range targets {
		for _, id := range ids {
			if strings.Contains(target.Ident, id) {
				filteredTargets = append(filteredTargets, target)
				break
			}
		}
	}

	return filteredTargets, nil
}
