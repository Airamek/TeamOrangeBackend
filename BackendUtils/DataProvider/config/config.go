package config

import (
	"fmt"
	"log"
)
import "os"
import "gopkg.in/yaml.v3"

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

type ConfigStructure struct {
	LdapSettings        *LdapData
	DataProviderBackend string
}

func GetConfig() *ConfigStructure {
	var err error
	var configS = new(ConfigStructure)
	configS.LdapSettings = new(LdapData)

	_, err = os.Stat("config.yaml")

	if err != nil {
		return CreateDefaultConf()
	}

	var data []byte
	data, err = os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("File exists, but it couldn't be opened")
	}
	yaml.Unmarshal(data, configS)

	return configS
}

func CreateDefaultConf() *ConfigStructure {
	var err error
	var configS = new(ConfigStructure)
	configS.LdapSettings = new(LdapData)

	fmt.Fprintf(os.Stderr, "config.yaml doesn't exist, creating it now\n")
	//top level
	configS.DataProviderBackend = "ldap"

	//ldap
	configS.LdapSettings.Url = "ldap://ldapserver"
	configS.LdapSettings.Port = 389
	configS.LdapSettings.Dn = "dc=example,dc=com"
	configS.LdapSettings.Binddn = "cn=Manager,dc=example,dc=com"
	configS.LdapSettings.Starttls = false
	configS.LdapSettings.Bindpass = "examplepass"
	configS.LdapSettings.Userlocation = "ou=Users,dc=example,dc=com"
	configS.LdapSettings.UserSearchAttribute = "cn"
	configS.LdapSettings.Grouplocation = "ou=Groups,dc=example,dc=com"
	configS.LdapSettings.UserIdentifierAttibute = "uid"
	configS.LdapSettings.UserNameAttribute = "cn"
	configS.LdapSettings.UserMailAttribute = "mail"
	configS.LdapSettings.UserMailAliasAttribute = "proxyAddresses"
	configS.LdapSettings.UserFilterClass = "person"
	configS.LdapSettings.GroupFilterClass = "group"

	//write out structure
	var file, _ = os.Create("config.yaml")
	var yamlData []byte
	yamlData, err = yaml.Marshal(configS)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	_, err = file.Write(yamlData)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return configS
}
