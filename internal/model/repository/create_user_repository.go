package repository

import (
	"context"
	"os"

	"github.com/HavocJean/study-go/internal/config/rest_error"
	"github.com/HavocJean/study-go/internal/logger"
	"github.com/HavocJean/study-go/internal/model"
)

const (
	MONGODB_DB = "MONGODB_DB"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_error.RestError) {
	logger.Info("init create user repository")

	collection_name := os.Getenv(MONGODB_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	value, err := userDomain.GetJSONValue()
	if err != nil {
		return nil, rest_error.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_error.NewInternalServerError(err.Error())
	}

	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil
}
