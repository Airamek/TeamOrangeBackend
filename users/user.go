package users

type User interface {
	GetName() string
	GetDisplayName() string
	GetMainEmail() string
	SetDisplayName()
	SetMainEmail()
	GetAliasEmails() []string
	AddAliasEmail(address string)
	DeleteAliasEmail(address string)
}
