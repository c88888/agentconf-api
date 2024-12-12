package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/n9e"
)

type CollectRuleService struct{}

// @function: GetCollectRuleById
// @description: 根据Id获取
// @param: info request.PageInfo
// @return: list interface{}, total int64, err error
func (c *CollectRuleService) GetCollectRuleById(id int) (cr model.CollectRule, err error) {
	err = global.MustGetGlobalDBByDBName("n9e").Where("id = ?", id).First(&cr).Error
	return
}

func (c *CollectRuleService) CreateCollectRule(cr model.CollectRule) (err error) {
	if cr.GroupId == 0 {
		return errors.New("业务组id不能为0或空")
	}
	db := global.MustGetGlobalDBByDBName("n9e")
	if !errors.Is(db.Where("name = ? and group_id = ?", cr.Name, cr.GroupId).First(&model.CollectRule{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("本业务组存在同名规则")
	}
	if err = db.Create(&cr).Error; err != nil {
		return
	}

	//if err := ClearCacheForUniqueIdents(context.Background(), cr.TargetIdents); err != nil {
	//	return fmt.Errorf("创建collectrule时，清理缓存失败: %v", err)
	//}

	return
}

// @function: GetCollectRuleList
// @description: 分页获取列表
// @param: info request.PageInfo
// @return: list interface{}, total int64, err error
func (b *CollectRuleService) GetCollectRuleList(cr model.CollectRule, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MustGetGlobalDBByDBName("n9e").Model(&model.CollectRule{}).Where("group_id=?", cr.GroupId)
	var crLists []model.CollectRule

	if cr.Name != "" {
		db = db.Where("name LIKE ?", "%"+cr.Name+"%")
	}

	err = db.Count(&total).Error

	if err != nil {
		return crLists, total, err
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
				return crLists, total, err
			}
			err = db.Order(OrderStr).Find(&crLists).Error
		} else {
			err = db.Order("id").Find(&crLists).Error
		}
	}

	// 用带索引的map来确保排序顺序
	resultMap := make(map[int]model.CollectRule)
	var mu sync.Mutex
	var wg sync.WaitGroup

	if len(crLists) >= 1 {
		for index, v := range crLists {
			wg.Add(1)
			go func(index int, v model.CollectRule) {
				defer wg.Done()
				mu.Lock()
				resultMap[index] = v
				mu.Unlock()
			}(index, v)
		}
	}

	wg.Wait()

	data := make([]model.CollectRule, len(resultMap))
	for i := 0; i < len(resultMap); i++ {
		data[i] = resultMap[i]
	}

	return data, total, err
}

// @function: UpdateCollectRule
// @description: 根据id更新CollectRule
// @param: cr model.CollectRule
// @return: error
func (c *CollectRuleService) UpdateCollectRule(cr model.CollectRule) error {
	if err := global.MustGetGlobalDBByDBName("n9e").
		Where("id = ?", cr.ID).
		First(&model.CollectRule{}).
		Updates(&cr).Error; err != nil {
		return err
	}

	//if err := ClearCacheForUniqueIdents(context.Background(), cr.TargetIdents); err != nil {
	//	return fmt.Errorf("修改collectrule时，清理缓存失败: %v", err)
	//}

	return nil
}

// @function: DeleteCollectRule
// @description: 删除CollectRule
// @param: cr model.CollectRule
// @return: error
func (c *CollectRuleService) DeleteCollectRule(cr model.CollectRule) error {
	if err := global.MustGetGlobalDBByDBName("n9e").Delete(&cr).Error; err != nil {
		return err
	}
	//if err := ClearCacheForUniqueIdents(context.Background(), cr.TargetIdents); err != nil {
	//	return fmt.Errorf("删除collectrule时，清理缓存失败: %v", err)
	//}
	return nil
}

// @function: SetCache
// @description: 缓存CollectRule
// @param: ctx context.Context, cacheKey, calculatedVersion string, configs map[string]map[string]*response.ConfigWithFormat
// @return: error
func SetCache(ctx context.Context, cacheKey, calculatedVersion string, configs []byte) error {
	if _, err := global.GVA_REDIS.Set(ctx, cacheKey+":version", calculatedVersion, time.Second*600).Result(); err != nil {
		return err
	}

	if _, err := global.GVA_REDIS.Set(ctx, cacheKey+":configs", configs, time.Second*600).Result(); err != nil {
		return err
	}

	return nil
}

func SetCollectRuleConfig(calculatedVersion string, configs []byte) error {
	db := global.MustGetGlobalDBByDBName("n9e")
	var existingConfig model.CollectRuleConfig

	err := db.Where("version = ?", calculatedVersion).First(&existingConfig).Error
	if err == nil {
		if existingConfig.Configs == string(configs) {
			return nil
		}

		err = db.Model(&existingConfig).Update("configs", string(configs)).Error
		if err != nil {
			return fmt.Errorf("failed to update configs for version %s: %v", calculatedVersion, err)
		}
		return nil
	}

	// 如果查询不到记录或发生其他错误，检查是否是 "record not found" 错误
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to query database: %v", err)
	}

	// 若没有该版本的记录，插入新记录
	err = db.Create(&model.CollectRuleConfig{
		Version: calculatedVersion,
		Configs: string(configs), // JSON 数据直接存储
	}).Error
	if err != nil {
		return fmt.Errorf("failed to save version and configs to database: %v", err)
	}

	return nil
}

// @function: ClearCache
// @description: 清除CollectRule缓存
// @param: ident string
// @return:
func ClearCache(ctx context.Context, cacheKey string) error {
	if err := global.GVA_REDIS.Del(ctx, cacheKey+":version").Err(); err != nil {
		return err
	}
	return global.GVA_REDIS.Del(ctx, cacheKey+":configs").Err()
}

// @function: ClearCacheForUniqueIdents
// @description: 根据CollectRule
// @param: ident string
// @return: list response.CollectRuleProviderResponse,err error
func ClearCacheForUniqueIdents(ctx context.Context, targetIdents string) error {
	identsSlice := strings.Split(targetIdents, " ")
	uniqueIdents := make(map[string]struct{})

	for _, ident := range identsSlice {
		if _, exists := uniqueIdents[ident]; !exists && ident != "" {
			uniqueIdents[ident] = struct{}{}
		}
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(uniqueIdents)) //接收所有错误
	for ident := range uniqueIdents {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			if err := ClearCache(ctx, v); err != nil {
				errChan <- fmt.Errorf("清理缓存失败key: %s, err: %v", v, err)
			}
		}(ident)
	}

	wg.Wait()
	close(errChan)

	var errList []string
	for err := range errChan {
		errList = append(errList, err.Error())
	}

	if len(errList) > 0 {
		return fmt.Errorf("清理缓存失败: %s", strings.Join(errList, "; "))
	}

	return nil
}
