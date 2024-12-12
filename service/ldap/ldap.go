package ldap

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	sysUserService "github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/go-ldap/ldap/v3"
	"github.com/gofrs/uuid/v5"
	"go.uber.org/zap"
)

type LdapService struct{}

func (*LdapService) GetAttributes(FieldMapping string) ([]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(FieldMapping), &m)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, v := range m {
		result = append(result, v)
	}
	return result, nil
}

func (*LdapService) GetMappings(FieldMapping string) (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(FieldMapping), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (ld *LdapService) LdapReq(user, pass string) (*ldap.SearchResult, error) {
	var conn *ldap.Conn
	var err error
	lc := global.GVA_CONFIG.LDAP
	addr := fmt.Sprintf("%s:%d", lc.LdapServer, lc.LdapPort)
	if lc.IsSSL {
		conn, err = ldap.DialTLS("tcp", addr, &tls.Config{InsecureSkipVerify: true})
	} else {
		conn, err = ldap.Dial("tcp", addr)
	}

	if err != nil {
		return nil, fmt.Errorf("ldap.error: cannot dial ldap(%s): %v", addr, err)
	}

	defer conn.Close()

	if !lc.IsSSL && lc.StartTLS {
		if err := conn.StartTLS(&tls.Config{InsecureSkipVerify: true}); err != nil {
			return nil, fmt.Errorf("ldap.error: conn startTLS fail: %v", err)
		}
	}

	// if bindUser is empty, anonymousSearch mode
	if lc.LdapDn != "" {
		// BindSearch mode
		if err := conn.Bind(lc.LdapDn, lc.LdapPassword); err != nil {
			return nil, fmt.Errorf("ldap.error: bind ldap fail: %v, use user(%s) to bind", err, lc.LdapOu)
		}
	}
	attributes, err := ld.GetAttributes(lc.FieldMapping)
	if err != nil {
		return nil, err
	}

	searchRequest := ldap.NewSearchRequest(
		lc.LdapOu, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, lc.PageLimit, lc.TimeOut, false,
		fmt.Sprintf(lc.LdapFilter, user), // The filter to apply
		attributes,
		nil,
	)

	sr, err := conn.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("ldap.error: ldap search fail: %v", err)
	}

	if len(sr.Entries) == 0 {
		return nil, fmt.Errorf("ldap.error: Username or password invalid")
	}

	if len(sr.Entries) > 1 {
		return nil, fmt.Errorf("ldap.error: search user(%s), multi entries found", user)
	}

	if err := conn.Bind(sr.Entries[0].DN, pass); err != nil {
		return nil, fmt.Errorf("Username or password invalid")
	}
	return sr, nil
}

// 三方登录
func (l *LdapService) FindUserByUserName(u *system.SysUser) (user *system.SysUser, err error) {
	var s system.SysUser
	err = global.GVA_DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&s).Error
	if err == nil { //跳过密码校验
		sysUserService.MenuServiceApp.UserAuthorityDefaultRouter(&s)
	}

	return &s, err
}

func (l *LdapService) LdapLogin(u *system.SysUser) (*system.SysUser, error) {
	sr, err := l.LdapReq(u.Username, u.Password)
	if err != nil {
		global.GVA_LOG.Error("ldap登陆失败!", zap.Error(err))
		return nil, err
	}

	user, err := l.FindUserByUserName(u)
	if err == nil { // is nil
		return user, nil
	}
	var userReturn system.SysUser
	if user.Username == "" {
		authority := system.SysAuthority{
			AuthorityId:   uint(1),
			ParentId:      utils.Pointer[uint](0),
			DefaultRouter: "dashboard",
			AuthorityName: "Common",
		}
		authorities := []system.SysAuthority{authority}
		pass := utils.RandomString(16)

		userData := &system.SysUser{
			UUID:        uuid.Must(uuid.NewV4()),
			Username:    u.Username,
			Password:    pass,
			NickName:    sr.Entries[0].GetAttributeValue("cn") + fmt.Sprintf(`(%s)`, u.Username),
			HeaderImg:   "",
			AuthorityId: 1,
			Authority:   authority,   //活跃角色
			Authorities: authorities, //可选角色
			Phone:       sr.Entries[0].GetAttributeValue("mobile"),
			Email:       sr.Entries[0].GetAttributeValue("mail"),
		}
		var us sysUserService.UserService
		userReturn, err = us.Register(*userData)
		if err != nil {
			global.GVA_LOG.Error("添加ldap用户失败!", zap.Error(err))
			return nil, err
		}
		return &userReturn, nil
	}
	return &userReturn, err
}
