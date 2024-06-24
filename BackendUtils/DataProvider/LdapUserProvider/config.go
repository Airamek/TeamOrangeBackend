package LdapUserProvider

import "main/BackendUtils/DataProvider/config"

type LdapData struct {
	Url                    string
	Port                   int
	Dn                     string
	Binddn                 string
	Bindpass               string
	Starttls               bool
	Userlocation           string
	UserFilterClass        string
	GroupFilterClass       string
	Grouplocation          string
	UserIdentifierAttibute string
	UserNameAttribute      string
	UserMailAttribute      string
	UserMailAliasAttribute string
	UserSearchAttribute    string
}

func InitConfig(name string) *LdapData {
	var conf *LdapData = new(LdapData)
	err := config.GetConfig(name, conf)
	if err != nil {
		val := config.CreateDefaultConf(name, CreateDefaultConfig())
		conf = val.(*LdapData)
	}
	return conf
}

func CreateDefaultConfig() *LdapData {
	LdapSettings := new(LdapData)
	LdapSettings.Url = "ldap://ldapserver"
	LdapSettings.Port = 389
	LdapSettings.Dn = "dc=example,dc=com"
	LdapSettings.Binddn = "cn=Manager,dc=example,dc=com"
	LdapSettings.Starttls = false
	LdapSettings.Bindpass = "examplepass"
	LdapSettings.Userlocation = "ou=Users,dc=example,dc=com"
	LdapSettings.UserSearchAttribute = "cn"
	LdapSettings.Grouplocation = "ou=Groups,dc=example,dc=com"
	LdapSettings.UserIdentifierAttibute = "uid"
	LdapSettings.UserNameAttribute = "cn"
	LdapSettings.UserMailAttribute = "mail"
	LdapSettings.UserMailAliasAttribute = "proxyAddresses"
	LdapSettings.UserFilterClass = "person"
	LdapSettings.GroupFilterClass = "group"
	return LdapSettings
}
