package users

var Sessions []Session

type Session struct {
	SessionID   string
	SessionUser *User
}
