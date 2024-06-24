package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"time"
)

type LdapUser struct {
	displayName          string
	displayNameCacheTime time.Time
	name                 string
	nameCacheTime        time.Time
	permLevel            string
	provider             *Provider
	mainMail             string
	mainMailCacheTime    time.Time
	aliasMails           []string
	aliasMailsCacheTime  time.Time
}

func (user LdapUser) checkIfPropertyExists(propertyName string) {

}

func (user LdapUser) setStringProperty(propertyName string, propertyPointer *string, cacheTime *time.Time) {

}

func (user LdapUser) getStringProperty(propertyName string, propertyPointer *string, cacheTime *time.Time) error {
	if time.Now().Sub(*cacheTime).Minutes() > 1 {
		searchRequest := ldap.NewSearchRequest(
			user.provider.Conf.UserIdentifierAttibute+"="+user.name+","+user.provider.Conf.Userlocation,
			ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
			fmt.Sprintf("(&(objectClass=%s))", user.provider.Conf.UserFilterClass),
			[]string{"dn", propertyName},
			nil,
		)
		searchResult, err := user.provider.Conn.Search(searchRequest)
		if err != nil {
			return err
		}
		for _, entry := range searchResult.Entries {
			*propertyPointer = entry.GetAttributeValue(propertyName)
		}

		*cacheTime = time.Now()
	}

	return nil
}

func (user LdapUser) getStringArrProperty(propertyName string, propertyPointer *[]string, cacheTime *time.Time) error {
	if time.Now().Sub(*cacheTime).Minutes() > 1 {
		searchRequest := ldap.NewSearchRequest(
			user.provider.Conf.UserIdentifierAttibute+"="+user.name+","+user.provider.Conf.Userlocation,
			ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
			fmt.Sprintf("(&(objectClass=%s))", user.provider.Conf.UserFilterClass),
			[]string{"dn", propertyName},
			nil,
		)
		searchResult, err := user.provider.Conn.Search(searchRequest)
		if err != nil {
			return err
		}
		for _, entry := range searchResult.Entries {
			*propertyPointer = entry.GetAttributeValues(propertyName)
		}

		*cacheTime = time.Now()
	}

	return nil
}

func (user LdapUser) GetDisplayName() string {
	user.getStringProperty(user.provider.Conf.UserNameAttribute, &user.displayName, &user.displayNameCacheTime)
	return user.displayName
}

func (user LdapUser) GetName() string {
	user.getStringProperty(user.provider.Conf.UserIdentifierAttibute, &user.name, &user.nameCacheTime)
	return user.name
}

func (user LdapUser) GetMainEmail() string {
	user.getStringProperty(user.provider.Conf.UserMailAttribute, &user.mainMail, &user.mainMailCacheTime)
	return user.mainMail
}

func (user LdapUser) SetMainEmail() {

}
func (user LdapUser) GetAliasEmails() []string {
	user.getStringArrProperty(user.provider.Conf.UserMailAliasAttribute, &user.aliasMails, &user.aliasMailsCacheTime)
	return user.aliasMails
}
func (user LdapUser) AddAliasEmail(address string) {

}
func (user LdapUser) DeleteAliasEmail(address string) {

}

func (user LdapUser) SetDisplayName() {

}
