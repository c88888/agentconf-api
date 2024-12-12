package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/n9e"
)

type PluginTplService struct{}

// @function: CreatePluginTpl
// @description: 创建PluginTpl
// @param: pt model.PluginTpl
// @return: err error
func (p *PluginTplService) CreatePluginTpl(pt model.PluginTpl) (err error) {
	if pt.Ident == "" {
		return errors.New("插件标识不能为空")
	}
	db := global.MustGetGlobalDBByDBName("n9e")
	if !errors.Is(db.Where("ident = ?", pt.Ident).First(&model.PluginTpl{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在同名插件")
	}
	if err = db.Create(&pt).Error; err != nil {
		return
	}
	return
}

// @function: UpdatePluginTpl
// @description: 根据id更新PluginTpl
// @param: pt model.PluginTpl
// @return: err error
func (p *PluginTplService) UpdatePluginTpl(pt model.PluginTpl) error {
	now := time.Now().Unix()
	pt.UpdatedAt = now
	return global.MustGetGlobalDBByDBName("n9e").
		Where("id = ?", pt.ID).
		First(&model.PluginTpl{}).
		Updates(&pt).Error
}

// @function: DeletePluginTpl
// @description: 删除PluginTpl
// @param: pt model.PluginTpl
// @return: error
func (*PluginTplService) DeletePluginTpl(pt model.PluginTpl) error {
	return global.MustGetGlobalDBByDBName("n9e").Delete(&pt).Error
}

// @function: GetPluginTplById
// @description: 根据Id获取
// @param: info request.PageInfo
// @return: list interface{}, total int64, err error
func (*PluginTplService) GetPluginTplById(id int64) (pt *model.PluginTpl, err error) {
	if id != 0 {
		err = global.MustGetGlobalDBByDBName("n9e").Where("id = ?", id).First(&pt).Error
		return
	}
	return nil, err
}

// @function: GetPluginTplByIdent
// @description: 根据Id获取
// @param: info request.PageInfo
// @return: list interface{}, total int64, err error
func (*PluginTplService) GetPluginTplByIdent(ident string) (pt *model.PluginTpl, err error) {
	if ident != "" {
		err = global.MustGetGlobalDBByDBName("n9e").Where("ident = ?", ident).First(&pt).Error
		return
	}
	return nil, err
}

// @function: GetBusiGroupList
// @description: 分页获取列表
// @param: info request.PageInfo
// @return: list interface{}, total int64, err error
func (*PluginTplService) GetPluginTplList(pt model.PluginTpl, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MustGetGlobalDBByDBName("n9e").Model(&model.PluginTpl{})
	var bgLists []model.PluginTpl

	if pt.Ident != "" {
		db = db.Where("ident LIKE ?", "%"+pt.Ident+"%")
	}
	if info.Keyword != "" {
		db = db.Where("ident LIKE ?", "%"+info.Keyword+"%")
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
			orderMap["ident"] = true

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
			err = db.Order("ident").Find(&bgLists).Error
		}
	}

	// 用带索引的map来确保排序顺序
	resultMap := make(map[int]model.PluginTpl)
	var mu sync.Mutex
	var wg sync.WaitGroup

	if len(bgLists) >= 1 {
		for index, v := range bgLists {
			wg.Add(1)
			go func(index int, v model.PluginTpl) {
				defer wg.Done()
				mu.Lock()
				resultMap[index] = v
				mu.Unlock()
			}(index, v)
		}
	}

	wg.Wait()

	data := make([]model.PluginTpl, len(resultMap))
	for i := 0; i < len(resultMap); i++ {
		data[i] = resultMap[i]
	}

	return data, total, err
}
