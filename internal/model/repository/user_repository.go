package repository

import (
	"github.com/HavocJean/study-go/internal/config/rest_error"
	"github.com/HavocJean/study-go/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_error.RestError)

	FindUserByEmail(email string) (model.UserDomainInterface, *rest_error.RestError)

	FindUserByID(id string) (model.UserDomainInterface, *rest_error.RestError)

	UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_error.RestError
}
