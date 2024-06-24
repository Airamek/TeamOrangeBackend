package DataProvider

import (
	"main/BackendUtils/DataProvider/LdapUserProvider"
	"main/BackendUtils/users"
)

type UserDataProvider interface {
	Init(name string)
	GetUsers() []users.User
	GetUsersData() []users.UserData
	AuthUser(username string, passwd string) users.User
}

func CreateProviderUser(providerType string) UserDataProvider {
	if providerType == "LdapUserProvider" {
		return new(LdapUserProvider.Provider)
	}
	return nil
}
