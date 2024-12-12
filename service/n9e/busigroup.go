package service

import (
	"fmt"
	"sync"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/n9e"
)

type BusiGroupService struct{}

func (*BusiGroupService) GetBusiGroupById(id int64) (bg *model.BusiGroup, err error) {
	if id != 0 {
		err = global.MustGetGlobalDBByDBName("n9e").Where("id = ?", id).First(&bg).Error
		return
	}
	return nil, err
}

// @function: GetBusiGroupList
// @description: 分页获取列表
// @param: info request.PageInfo
// @return: list interface{}, total int64, err error
func (*BusiGroupService) GetBusiGroupList(bg model.BusiGroup, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MustGetGlobalDBByDBName("n9e").Model(&model.BusiGroup{})
	var bgLists []model.BusiGroup

	if bg.Name != "" {
		db = db.Where("name LIKE ?", "%"+bg.Name+"%")
	}

	err = db.Count(&total).Error

	if err != nil {
		return bgLists, total, err
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
				return bgLists, total, err
			}
			err = db.Order(OrderStr).Find(&bgLists).Error
		} else {
			err = db.Order("name").Find(&bgLists).Error
		}
	}

	// 用带索引的map来确保排序顺序
	resultMap := make(map[int]model.BusiGroup)
	var mu sync.Mutex
	var wg sync.WaitGroup

	if len(bgLists) >= 1 {
		for index, v := range bgLists {
			wg.Add(1)
			go func(index int, v model.BusiGroup) {
				defer wg.Done()
				mu.Lock()
				resultMap[index] = v
				mu.Unlock()
			}(index, v)
		}
	}

	wg.Wait()

	data := make([]model.BusiGroup, len(resultMap))
	for i := 0; i < len(resultMap); i++ {
		data[i] = resultMap[i]
	}

	return data, total, err
}
