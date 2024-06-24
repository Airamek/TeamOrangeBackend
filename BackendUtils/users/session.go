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

func AddSession(sessionUser *User) *Session {
	session := new(Session)
	session.User = sessionUser
	session.Token.Token = uniuri.NewLen(32)
	session.Token.expirationTime = time.Now().Add(time.Hour * time.Duration(24*7))
	session.Token.invalid = false
	Sessions = append(Sessions, *session)
	return session
}

func CheckSession(token string) *Session {
	for _, session := range Sessions {
		if session.Token.Token == token && session.Token.expirationTime.After(time.Now()) && !session.Token.invalid {
			return &session
		}
	}
	return nil
}
