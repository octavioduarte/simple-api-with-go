package domain

import (
	"log"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	Email    string `json:"email" gorm:"type:varchar(255);unique_index"`
	Name     string `json:"name"  gorm:"type:varchar(255)"`
	Password string `json:"-"     gorm:"type:varchar(255);unique_index"`
	Token    string `json:"token" gorm:"type:varchar(255)"`
}

func NewUser() *User {
	return &User{}
}

func (user *User) Prepare() error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("Error during the password generation: %v", err)
	}

	user.Password = string(password)
	user.Token = uuid.NewV4().String()

	return err
}
