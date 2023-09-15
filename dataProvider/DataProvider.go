package dataProvider

import (
	"main/dataProvider/config"
	"main/dataProvider/ldap"
	"main/users"
)

type DataProvider interface {
	Init()
	GetUsers() []users.User
	AuthUser(username string, passwd string) bool
}

func CreateProvider() DataProvider {
	if config.GetConfig().DataProviderBackend == "ldap" {
		return new(ldap.Provider)
	}
	return nil
}
