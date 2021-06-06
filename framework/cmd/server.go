package main

import (
	"fmt"
	"log"

	"github.com/octavioduarte/simple-api-with-go/application/repositories"
	"github.com/octavioduarte/simple-api-with-go/domain"
	"github.com/octavioduarte/simple-api-with-go/framework/utils"
)

func main() {

	db := utils.ConnectDB()

	user := domain.User{
		Name:     "Fake User",
		Email:    "fake@server.com",
		Password: "secret_123",
	}

	userRepo := repositories.UserRepositoryDB{Db: db}
	result, err := userRepo.Insert(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.Name, "registerd with success !")
}
