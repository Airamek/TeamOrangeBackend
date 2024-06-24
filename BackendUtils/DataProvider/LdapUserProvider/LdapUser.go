package LdapUserProvider

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"log"
	"main/BackendUtils/users"
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

func (user LdapUser) checkIfPropertyExists(propertyName string) bool {

	searchResult, err := performSearch(&user, propertyName, nil)
	if err != nil {
		log.Fatal(err)
	}
	return len(searchResult.Entries) > 0
}

func (user *LdapUser) setStringProperty(propertyName string, propertyPointer *string, cacheTime *time.Time) error {
	searchResult, err := performSearch(user, propertyName, cacheTime)
	if err != nil {
		return err
	}
	if searchResult != nil {
		for _, entry := range searchResult.Entries {
			*propertyPointer = entry.GetAttributeValue(propertyName)
		}
	}
	return nil
}

func (user *LdapUser) getStringProperty(propertyName string, propertyPointer *string, cacheTime *time.Time) error {
	if time.Since(*cacheTime).Minutes() <= 1 {
		return nil // Cache is valid
	}
	err := user.setStringProperty(propertyName, propertyPointer, cacheTime)
	return err
}

func (user *LdapUser) getStringArrProperty(propertyName string, propertyPointer *[]string, cacheTime *time.Time) error {
	if time.Since(*cacheTime).Minutes() <= 1 {
		return nil // Cache is valid
	}
	searchResult, err := performSearch(user, propertyName, cacheTime)
	if err != nil {
		return err
	}
	if searchResult != nil {
		for _, entry := range searchResult.Entries {
			*propertyPointer = entry.GetAttributeValues(propertyName)
		}
	}
	return nil
}

func (user *LdapUser) GetDisplayName() string {
	user.getStringProperty(user.provider.Conf.UserNameAttribute, &user.displayName, &user.displayNameCacheTime)
	return user.displayName
}

func (user *LdapUser) GetName() string {
	return user.name
}

func (user *LdapUser) GetMainEmail() string {
	user.getStringProperty(user.provider.Conf.UserMailAttribute, &user.mainMail, &user.mainMailCacheTime)
	return user.mainMail
}

func (user *LdapUser) SetMainEmail() {
	// Implementation for setting main email
}

func (user *LdapUser) GetAliasEmails() []string {
	user.getStringArrProperty(user.provider.Conf.UserMailAliasAttribute, &user.aliasMails, &user.aliasMailsCacheTime)
	return user.aliasMails
}

func (user *LdapUser) AddAliasEmail(address string) {
	// Implementation for adding alias email
}

func (user *LdapUser) DeleteAliasEmail(address string) {
	// Implementation for deleting alias email
}

func (user *LdapUser) SetDisplayName() {
	// Implementation for setting display name
}

func (user *LdapUser) GetData() users.UserData {
	data := users.UserData{
		Name:        user.GetName(),
		DisplayName: user.GetDisplayName(),
		MainEmail:   user.GetMainEmail(),
		AliasEmails: user.GetAliasEmails(),
	}
	return data
}

func performSearch(user *LdapUser, propertyName string, cacheTime *time.Time) (*ldap.SearchResult, error) {
	searchRequest := ldap.NewSearchRequest(
		fmt.Sprintf("%s=%s,%s", user.provider.Conf.UserSearchAttribute, user.name, user.provider.Conf.Userlocation),
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=%s))", user.provider.Conf.UserFilterClass),
		[]string{"dn", propertyName},
		nil,
	)
	searchResult, err := user.provider.Conn.Search(searchRequest)
	if err != nil {
		return nil, err
	}
	if cacheTime != nil {
		*cacheTime = time.Now()
	}
	return searchResult, nil
}
