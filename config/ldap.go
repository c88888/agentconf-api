package config

import (
	"encoding/json"
)

type Ldap struct {
	Enable       bool   `mapstructure:"enable" json:"enable" yaml:"enable"`                      // 是否启用
	LdapServer   string `mapstructure:"ldap-server" json:"ldap-server" yaml:"ldap-server"`       // ldap地址
	LdapPort     int    `mapstructure:"ldap-port" json:"ldap-port" yaml:"ldap-port"`             // ldap端口
	IsSSL        bool   `mapstructure:"is-ssl" json:"is-ssl" yaml:"is-ssl"`                      // 是否SSL
	StartTLS     bool   `mapstructure:"start-tls" json:"start-tls" yaml:"start-tls"`             // 是否TLS
	LdapDn       string `mapstructure:"ldap-dn" json:"ldap-dn" yaml:"ldap-dn"`                   // 绑定DN
	LdapPassword string `mapstructure:"ldap-password" json:"ldap-password" yaml:"ldap-password"` // 密码
	LdapOu       string `mapstructure:"ldap-ou" json:"ldap-ou" yaml:"ldap-ou"`                   // 用户过滤OU
	LdapFilter   string `mapstructure:"ldap-filter" json:"ldap-filter" yaml:"ldap-filter"`       //用户过滤规则
	FieldMapping string `mapstructure:"field-mapping" json:"field-mapping" yaml:"field-mapping"` //用户属性映射
	TimeOut      int    `mapstructure:"timeout" json:"timeout" yaml:"timeout"`                   //连接超时时间
	PageLimit    int    `mapstructure:"pagelimit" json:"pagelimit" yaml:"pagelimit"`             //搜索分页数量
}

func (l *Ldap) GetAttributes() ([]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(l.FieldMapping), &m)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, v := range m {
		result = append(result, v)
	}
	return result, nil
}

func (l *Ldap) GetMappings() (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(l.FieldMapping), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
