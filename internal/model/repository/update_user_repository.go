package repository

import (
	"context"
	"os"

	"github.com/HavocJean/study-go/internal/config/rest_error"
	"github.com/HavocJean/study-go/internal/logger"
	"github.com/HavocJean/study-go/internal/model"
	"github.com/HavocJean/study-go/internal/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_error.RestError {
	logger.Info("init createUser repository", zap.String("journey", "updateUser"))

	collection_name := os.Getenv(MONGODB_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)
	userIdHex, err := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error trying to update user", err, zap.String("journey", "updateUser"))
		return rest_error.NewInternalServerError(err.Error())
	}

	logger.Info(
		"updateUser repository successfully",
		zap.String("userId", value.ID.Hex()),
		zap.String("journey", "updateUser"),
	)

	return nil
}
