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
	GetData() UserData
}

type UserData struct {
	Name        string
	DisplayName string
	MainEmail   string
	AliasEmails []string
}
