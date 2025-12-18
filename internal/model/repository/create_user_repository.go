package repository

import (
	"context"
	"os"

	"github.com/HavocJean/study-go/internal/config/rest_error"
	"github.com/HavocJean/study-go/internal/logger"
	"github.com/HavocJean/study-go/internal/model"
	"github.com/HavocJean/study-go/internal/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

const (
	MONGODB_DB = "MONGODB_DB"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_error.RestError) {
	logger.Info("init createUser repository", zap.String("journey", "createUser"))

	collection_name := os.Getenv(MONGODB_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to create user", err, zap.String("journey", "createUser"))
		return nil, rest_error.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	logger.Info(
		"CreateUser repository successfully",
		zap.String("userId", value.ID.Hex()),
		zap.String("journey", "createUser"),
	)

	return converter.ConvertEntityToDomain(*value), nil
}
