package config

import (
	"log"
)
import "os"
import "fmt"
import "gopkg.in/yaml.v3"

type LdapData struct {
	Url          string
	Port         string
	Dn           string
	Binddn       string
	Bindpass     string
	Starttls     bool
	Userlocation string
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
		fmt.Fprintf(os.Stderr, "config.yaml doesn't exist, creating it now\n")
		configS.DataProviderBackend = "ldap"
		configS.LdapSettings.Url = "ldap://ldapserver"
		configS.LdapSettings.Port = "389"
		configS.LdapSettings.Dn = "dc=example,dc=com"
		configS.LdapSettings.Binddn = "cn=Manager,dc=example,dc=com"
		configS.LdapSettings.Starttls = false
		configS.LdapSettings.Bindpass = "examplepass"
		configS.LdapSettings.Userlocation = "ou=Users,dc=example,dc=com"
		var file, _ = os.Create("config.yaml")
		var yamlData []byte
		yamlData, err = yaml.Marshal(configS)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		file.Write(yamlData)
		return configS
	}

	var data []byte
	data, err = os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("File exist, but it couldn't be opened")
	}
	yaml.Unmarshal(data, configS)

	return configS
}
