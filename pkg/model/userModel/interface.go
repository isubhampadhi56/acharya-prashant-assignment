package usermodel

type UserSignUp interface {
	SetPassword(string) error
	Save() error
}

type UserLogin interface {
	ValidatePassword(string) error
}
