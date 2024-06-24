package LdapUserProvider

import (
	"crypto/tls"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"log"
	"main/BackendUtils/users"
	"time"
)

type Provider struct {
	Conf        *LdapData
	Conn        *ldap.Conn
	UserArr     []users.User
	userArrTime time.Time
}

func (connData *Provider) Init(name string) {
	connData.Conf = InitConfig(name)
	// The username and password we want to check
	var err error
	fmt.Println(connData.Conf.Url)
	connData.Conn, err = ldap.DialURL(fmt.Sprintf("%s:%d", connData.Conf.Url, connData.Conf.Port))
	if err != nil {
		log.Fatal(err)
	}

	if connData.Conf.Starttls {
		//defer connData.Conn.Close()

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
	println(time.Now().Sub(connData.userArrTime).Minutes())
	if time.Now().Sub(connData.userArrTime).Minutes() > 1 {
		searchRequest := ldap.NewSearchRequest(
			connData.Conf.Userlocation,
			ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
			fmt.Sprintf("(&(objectClass=%s))", connData.Conf.UserFilterClass),
			[]string{"dn", connData.Conf.UserIdentifierAttibute},
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
		connData.userArrTime = time.Now()
		connData.UserArr = userArr
		return userArr
	}
	return connData.UserArr
}

func (connData *Provider) GetUsersData() []users.UserData {
	var usersData []users.UserData
	users := connData.GetUsers()
	for _, user := range users {
		usersData = append(usersData, user.GetData())
	}
	return usersData
}

func (connData *Provider) AuthUser(username string, passwd string) users.User {
	//searchRequest := LdapUserProvider.NewSearchRequest(
	//	fmt.Sprintf("uid=%s,%s", username, connData.Conf.Userlocation), // The base dn to search
	//	LdapUserProvider.ScopeWholeSubtree, LdapUserProvider.NeverDerefAliases, 0, 0, false,
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
	var err = connData.Conn.Bind(fmt.Sprintf("%s=%s,%s", connData.Conf.UserSearchAttribute, username, connData.Conf.Userlocation), passwd)
	if err != nil {
		log.Print(err)
		return nil
	}
	user := new(LdapUser)
	user.name = username
	user.permLevel = "user"
	user.provider = connData
	return user
}
