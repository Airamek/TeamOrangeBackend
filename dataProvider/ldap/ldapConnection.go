package ldap

import (
	"crypto/tls"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"log"
	"main/dataProvider/config"
	"main/users"
)

type Provider struct {
	Conf *config.LdapData
	Conn *ldap.Conn
}

func (connData *Provider) Init() {
	connData.Conf = config.GetConfig().LdapSettings
	// The username and password we want to check
	var err error

	connData.Conn, err = ldap.DialURL(connData.Conf.Url + ":" + connData.Conf.Port)
	if err != nil {
		log.Fatal(err)
	}

	if connData.Conf.Starttls {
		defer connData.Conn.Close()

		// Reconnect with TLS
		err = connData.Conn.StartTLS(&tls.Config{InsecureSkipVerify: true})
		if err != nil {
			log.Fatal(err)
		}

	}

	// Bind with admin user
	err = connData.Conn.Bind(connData.Conf.Binddn, connData.Conf.Bindpass)
	if err != nil {
		log.Fatal(err)
	}

}

func (connData *Provider) GetUsers() []users.User {
	// Search
	searchRequest := ldap.NewSearchRequest(
		connData.Conf.Userlocation,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=%s))", connData.Conf.UserFilterClass),
		[]string{"dn", connData.Conf.UserNameAttribute},
		nil,
	)
	searchResult, err := connData.Conn.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	var userArr []users.User

	for _, entry := range searchResult.Entries {
		user := new(LdapUser)
		user.name = entry.GetAttributeValue(connData.Conf.UserIdentifierAttibute)
		user.permLevel = "user"
		user.provider = connData
		userArr = append(userArr, user)
	}

	return userArr
}

func (connData *Provider) AuthUser(username string, passwd string) bool {
	//searchRequest := ldap.NewSearchRequest(
	//	fmt.Sprintf("uid=%s,%s", username, connData.Conf.Userlocation), // The base dn to search
	//	ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
	//	"(&(objectClass=organizationalPerson))", // The filter to apply
	//	[]string{"dn", "userPassword"},          // A list attributes to retrieve
	//	nil,
	//)
	//searchResult, err := connData.Conn.Search(searchRequest)
	//if err != nil {
	//	return false
	//}
	//encoder := SSHAEncoder{}
	//return encoder.Matches([]byte(searchResult.Entries[0].GetAttributeValue("userPassword")), []byte(passwd))
	var err = connData.Conn.Bind(username, passwd)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
