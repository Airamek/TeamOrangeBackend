package users

import (
	"github.com/dchest/uniuri"
	"time"
)

var Sessions []Session

type Session struct {
	Token Token
	User  *User
}

func addSession(sessionUser *User) {
	session := new(Session)
	session.User = sessionUser
	session.Token.Token = uniuri.NewLen(32)
	session.Token.expirationTime = time.Now().Add(time.Hour * time.Duration(24*7))
}
