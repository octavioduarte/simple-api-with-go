package domain

import (
	"log"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	Email    string
	Name     string
	Password string
	Token    string
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
