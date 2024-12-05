package usermodel

import (
	"fmt"
	"log"

	"github.com/api-assignment/pkg/utils/db"
	"golang.org/x/crypto/bcrypt"
)

type UserData struct {
	Id       uint64 `gorm:"primaryKey,autoIncrement" json:"userId" validate:"required"`
	Email    string `gorm:"unique;not null" json:"email" validate:"required,email"`
	Password string `gorm:"not null" json:"-" validate:"required"`
}

func (user *UserData) SetPassword(plainPassword string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *UserData) ValidatePassword(plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plainPassword))
	return err
}

func (user *UserData) Save() error {
	dbConn, err := db.GetDBConn()
	dbConn.AutoMigrate(&UserData{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(user)
	result := dbConn.GetDB().Create(user)
	return result.Error
}

func CreateUser(email string) UserSignUp {
	return &UserData{
		Email: email,
	}
}

func FindUserByEmail(email string) (UserLogin, error) {
	var user UserData
	dbConn, err := db.GetDBConn()
	if err != nil {
		log.Println(err)
	}
	result := dbConn.GetDB().Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
