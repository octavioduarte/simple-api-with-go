package repositories

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/octavioduarte/simple-api-with-go/domain"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}

type UserRepositoryDB struct {
	Db *gorm.DB
}

func (repo UserRepositoryDB) Insert(user *domain.User) (*domain.User, error) {

	err := user.Prepare()

	if err != nil {
		log.Fatalf("Ocurred a error while prepare user data: %v", err)
		return user, err
	}

	err = repo.Db.Create(user).Error

	if err != nil {
		log.Fatalf("Ocurred a error while save user on db: %v", err)
		return user, err
	}

	return user, nil
}
