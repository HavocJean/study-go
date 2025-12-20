package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/HavocJean/study-go/internal/config/rest_error"
	"github.com/HavocJean/study-go/internal/logger"
	"github.com/HavocJean/study-go/internal/model"
	"github.com/HavocJean/study-go/internal/model/repository/entity"
	"github.com/HavocJean/study-go/internal/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (u *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_error.RestError) {
	logger.Info("Find user by email repository", zap.String("journey", "FindUserByEmail"))

	collection_name := os.Getenv(MONGODB_DB)
	collection := u.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)
			logger.Error(errorMessage, err, zap.String("journey", "FindUserByEmail"))
			return nil, rest_error.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("journey", "FindUserByEmail"))
		return nil, rest_error.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmail executed successfully", zap.String("journey", "FindUserByEmail"))

	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (u *userRepository) FindUserByID(id string) (model.UserDomainInterface, *rest_error.RestError) {
	logger.Info("Find user by id repository", zap.String("journey", "FindUserByID"))

	collection_name := os.Getenv(MONGODB_DB)
	collection := u.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this id: %s", id)
			logger.Error(errorMessage, err, zap.String("journey", "FindUserByid"))
			return nil, rest_error.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by id"
		logger.Error(errorMessage, err, zap.String("journey", "FindUserByID"))
		return nil, rest_error.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByID executed successfully", zap.String("journey", "FindUserByID"))

	return converter.ConvertEntityToDomain(*userEntity), nil
}
