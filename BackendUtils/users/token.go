package users

import "time"

type Token struct {
	Token          string
	expirationTime time.Time
	invalid        bool
}
