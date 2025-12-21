package repository

import (
	"context"
	"os"

	"github.com/HavocJean/study-go/internal/config/rest_error"
	"github.com/HavocJean/study-go/internal/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(userId string) *rest_error.RestError {
	logger.Info("init DeleteUser repository", zap.String("journey", "DeleteUser"))

	collection_name := os.Getenv(MONGODB_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userIdHex, err := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error trying to Delete user", err, zap.String("journey", "DeleteUser"))
		return rest_error.NewInternalServerError(err.Error())
	}

	logger.Info(
		"DeleteUser repository successfully",
		zap.String("userId", userIdHex.Hex()),
		zap.String("journey", "DeleteUser"),
	)

	return nil
}
