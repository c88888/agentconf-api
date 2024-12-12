package service

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/n9e"
	"github.com/flipped-aurora/gin-vue-admin/server/model/n9e/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type ConfigService struct{}

//	func (c *ConfigService) CollectRuleProvide(version, ident string) (*response.CollectRuleProviderResponse, error) {
//		db := global.MustGetGlobalDBByDBName("n9e")
//		var rules []model.CollectRule
//		err := db.Where("FIND_IN_SET(?, REPLACE(target_idents,',',' ')) > 0 AND disabled != ?", ident, 1).Find(&rules).Error
//		if err != nil {
//			return nil, err
//		}
//
//		configs := make(map[string]*response.ConfigWithFormat)
//		for _, rule := range rules {
//			configs[rule.Type] = &response.ConfigWithFormat{
//				Config: rule.Data,
//				Format: "toml",
//			}
//		}
//
//		if len(configs) < 1 {
//			return &response.CollectRuleProviderResponse{
//				Version: "",
//				Configs: nil,
//			}, nil
//		}
//		calCfg, err := json.Marshal(configs)
//		if err != nil {
//			return nil, err
//		}
//		calculatedVersion := utils.MD5V(calCfg)
//		return &response.CollectRuleProviderResponse{
//			Version: calculatedVersion,
//			Configs: configs,
//		}, nil
//	}
func (c *ConfigService) checkSum(k string) string {
	chkSum, err := json.Marshal(k)
	if err != nil {
		return ""
	}
	return utils.MD5V(chkSum)
}

func (c *ConfigService) CollectRuleProvide(version, ident string) (*response.CollectRuleProviderResponse, error) {
	db := global.MustGetGlobalDBByDBName("n9e")
	var rules []model.CollectRule
	err := db.Where("FIND_IN_SET(?, REPLACE(target_idents,' ',',')) > 0 AND disabled != ?", ident, 1).Find(&rules).Error
	if err != nil {
		return nil, err
	}

	configs := make(map[string]map[string]*response.ConfigWithFormat)
	for _, rule := range rules {
		if _, exists := configs[rule.Type]; !exists {
			configs[rule.Type] = make(map[string]*response.ConfigWithFormat)
		}
		sum := c.checkSum(rule.Data)
		if c.checkSum(rule.Data) != "" {
			configs[rule.Type][rule.Name+sum] = &response.ConfigWithFormat{
				Config: rule.Data,
				Format: "toml",
			}
		}
	}

	if configs == nil || len(configs) < 1 {
		return &response.CollectRuleProviderResponse{
			Version: "",
			Configs: nil,
		}, nil
	}
	calCfg, err := json.Marshal(configs)
	if err != nil {
		return nil, err
	}
	calculatedVersion := utils.MD5V(calCfg)
	return &response.CollectRuleProviderResponse{
		Version: calculatedVersion,
		Configs: configs,
	}, nil
}

// @function: CollectRuleProvide
// @description: agent获取采集规则
// @param: ident string
// @return: *response.CollectRuleProviderResponse,error
//func (c *ConfigService) CollectRuleProvide1(version, ident string) (*response.CollectRuleProviderResponse, error) {
//	cacheKey := fmt.Sprintf("config:%s", ident)
//
//	var ruleConfig model.CollectRuleConfig
//	db := global.MustGetGlobalDBByDBName("n9e")
//	err := db.Where("version = ?", version).First(&ruleConfig).Error
//	if err == nil {
//		var configs map[string]map[string]*response.ConfigWithFormat
//		err = json.Unmarshal([]byte(ruleConfig.Configs), &configs)
//		if err != nil {
//			return nil, fmt.Errorf("failed to unmarshal configs from database: %v", err)
//		}
//
//		err = SetCache(context.Background(), cacheKey, version, []byte(ruleConfig.Configs))
//		if err != nil {
//			return nil, err
//		}
//
//		return &response.CollectRuleProviderResponse{
//			Version: ruleConfig.Version,
//			Configs: configs,
//		}, nil
//	}
//
//	// 生成新配置写入redis和数据库
//	var rules []model.CollectRule
//	err = db.Where("target_idents LIKE ?", "%"+ident+"%").Find(&rules).Error
//	if err != nil {
//		return nil, err
//	}
//
//	configs := make(map[string]map[string]*response.ConfigWithFormat)
//	for _, rule := range rules {
//		if _, exists := configs[rule.Type]; !exists {
//			configs[rule.Type] = make(map[string]*response.ConfigWithFormat)
//		}
//		configs[rule.Type][rule.Name] = &response.ConfigWithFormat{
//			Config: rule.Data,
//			Format: "toml",
//		}
//	}
//
//	//序列化configs,确保calculatedVersion一致性
//	calCfg, err := json.Marshal(configs)
//	if err != nil {
//		return nil, err
//	}
//	calculatedVersion := utils.MD5V(calCfg)
//
//	err = SetCache(context.Background(), cacheKey, calculatedVersion, calCfg)
//	if err != nil {
//		return nil, err
//	}
//	err = SetCollectRuleConfig(calculatedVersion, calCfg) //写库
//	if err != nil {
//		return nil, err
//	}
//	return &response.CollectRuleProviderResponse{
//		Version: calculatedVersion,
//		Configs: configs,
//	}, nil
//}

//func (*ConfigService) CheckCache(version, agentHostname string) (*response.CollectRuleProviderResponse, error) {
//	cacheKey := fmt.Sprintf("config:%s", agentHostname)
//
//	cachedVersion, err := global.GVA_REDIS.Get(context.Background(), cacheKey+":version").Result()
//	if err == redis.Nil { //未命中version
//		return nil, nil
//	}
//	if err != nil { //其它Redis错误
//		return nil, err
//	}
//
//	if version != cachedVersion { //version不匹配，返回nil
//		return nil, nil
//	}
//
//	cachedConfigs, err := global.GVA_REDIS.Get(context.Background(), cacheKey+":configs").Bytes()
//	if err == redis.Nil { //未命中configs
//		return nil, nil
//	}
//	if err != nil { //其它Redis错误
//		return nil, err
//	}
//
//	// 反序列化后返回version和configs
//	var configs map[string]map[string]*response.ConfigWithFormat
//	if err = json.Unmarshal(cachedConfigs, &configs); err != nil {
//		return nil, err
//	}
//
//	resp := &response.CollectRuleProviderResponse{
//		Version: cachedVersion,
//		Configs: configs,
//	}
//
//	return resp, nil
//}
