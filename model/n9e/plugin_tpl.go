package model

type PluginTpl struct {
	ID    int64  `gorm:"primaryKey;autoIncrement" json:"id"`   // 唯一标识符
	Ident string `gorm:"size:191;not null;index" json:"ident"` // plugin_tpl 的标识符
	//Logo      string `gorm:"size:191;not null" json:"logo"`                  // plugin_tpl 的 logo
	Toml      string `gorm:"type:text;not null" json:"toml"`                 // plugin_tpl 的 TOML 内容
	Readme    string `gorm:"type:text;not null" json:"readme"`               // plugin_tpl 的说明文档
	Note      string `gorm:"size:4096;not null" json:"note"`                 // plugin_tpl 的描述
	CreatedAt int64  `gorm:"not null;default:0" json:"created_at"`           // 创建时间
	CreatedBy string `gorm:"size:191;not null;default:''" json:"created_by"` // 创建者
	UpdatedAt int64  `gorm:"not null;default:0" json:"updated_at"`           // 更新时间
	UpdatedBy string `gorm:"size:191;not null;default:''" json:"updated_by"` // 更新者
}

func (*PluginTpl) TableName() string {
	return "plugin_tpl"
}

/*
CREATE TABLE `plugin_tpl` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '''unique identifier''',
  `ident` varchar(191) NOT NULL COMMENT '''identifier of plugin_tpl''',
  `toml` text NOT NULL COMMENT '''toml of plugin_tpl''',
  `readme` text NOT NULL COMMENT '''readme of plugin_tpl''',
  `note` varchar(4096) NOT NULL COMMENT '''description of plugin_tpl''',
  `created_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '''create time''',
  `created_by` varchar(191) NOT NULL DEFAULT '' COMMENT '''creator''',
  `updated_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '''update time''',
  `updated_by` varchar(191) NOT NULL DEFAULT '' COMMENT '''updater''',
  PRIMARY KEY (`id`),
  KEY `idx_ident` (`ident`)
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8mb4
*/
