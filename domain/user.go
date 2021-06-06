package domain

type User struct {
	Base
	Email    string
	Name     string
	Password string
	Token    string
}
