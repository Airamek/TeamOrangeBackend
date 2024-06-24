package DataProvider

import (
	"main/BackendUtils/DataProvider/config"
	"main/BackendUtils/DataProvider/ldap"
	"main/BackendUtils/users"
)

type DataProvider interface {
	Init()
	GetUsers() []users.User
	GetUsersData() []users.UserData
	AuthUser(username string, passwd string) users.User
}

func CreateProvider() DataProvider {
	if config.GetConfig().DataProviderBackend == "ldap" {
		return new(ldap.Provider)
	}
	return nil
}
