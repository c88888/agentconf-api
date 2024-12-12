package model

type CollectRule struct {
	ID           int64  `json:"id" gorm:"id"`
	GroupId      int64  `json:"group_id" gorm:"group_id"` // busi group id
	Cluster      string `json:"cluster" gorm:"cluster"`
	TargetIdents string `json:"target_idents" gorm:"type:text"` // ident list, split by space

	TargetTags string `json:"target_tags" gorm:"target_tags"` // filter targets by tags, split by space
	Name       string `json:"name" gorm:"name"`
	Note       string `json:"note" gorm:"note"`
	Step       int64  `json:"step" gorm:"step"`
	Type       string `json:"type" gorm:"type"` // e.g. port proc log plugin
	Data       string `json:"data" gorm:"data"`
	AppendTags string `json:"append_tags" gorm:"append_tags"` // split by space: e.g. mod=n9e dept=cloud
	CreateAt   int64  `json:"create_at" gorm:"create_at"`
	CreateBy   string `json:"create_by" gorm:"create_by"`
	UpdateAt   int64  `json:"update_at" gorm:"update_at"`
	UpdateBy   string `json:"update_by" gorm:"update_by"`

	Disabled *bool `json:"disabled" gorm:"column:disabled"` //'0:enabled 1:disabled'
}

func (*CollectRule) TableName() string {
	return "collect_rule"
}

type CollectRuleConfig struct {
	ID      uint64 `gorm:"primaryKey;autoIncrement" json:"id"`     // 主键
	Version string `gorm:"size:64;not null;unique" json:"version"` // 版本号，用于校验配置是否发生变化
	Configs string `gorm:"type:text;not null" json:"configs"`      // 配置内容，JSON 格式
	GroupId int64  `json:"group_id" gorm:"group_id"`               // busi group id

	CreateAt int64  `json:"create_at" gorm:"create_at"`
	CreateBy string `json:"create_by" gorm:"create_by"`
	UpdateAt int64  `json:"update_at" gorm:"update_at"`
	UpdateBy string `json:"update_by" gorm:"update_by"`
}

func (*CollectRuleConfig) TableName() string {
	return "collect_rule_config"
}

/*
CREATE TABLE `collect_rule_config` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `group_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'busi group id',
`version` VARCHAR(64) NOT NULL,
`configs` TEXT NOT NULL,
  `append_tags` varchar(255) NOT NULL DEFAULT '' COMMENT 'split by space: e.g. mod=n9e dept=cloud',
  `create_at` bigint(20) NOT NULL DEFAULT '0',
  `create_by` varchar(64) NOT NULL DEFAULT '',
  `update_at` bigint(20) NOT NULL DEFAULT '0',
  `update_by` varchar(64) NOT NULL DEFAULT '',
);
*/
