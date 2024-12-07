package usermodel

import "time"

type UserSignUp interface {
	SetPassword(string) error
	Save() error
}

type UserLogin interface {
	ValidatePassword(string) error
	GetUserID() uint64
	GetUserStatus() bool
	GetUserLastUpdated() time.Time
}

type UserStatus interface {
	GetUserStatus() bool
	Disable() error
	Enable() error
	Save() error
}
