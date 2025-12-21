package controller

import (
	"net/http"

	"github.com/HavocJean/study-go/internal/config/rest_error"
	"github.com/HavocJean/study-go/internal/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser controller", zap.String("journey", "DeleteUser"))

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_error.NewBadRequestError("Invalid userid, must be a hex value")
		c.JSON(int(errRest.Code), errRest)
		return
	}

	err := uc.service.DeleteUserServices(userId)
	if err != nil {
		logger.Error("Error try to call DeleteUser service", err, zap.String("journey", "DeleteUser"))

		c.JSON(int(err.Code), err)
		return
	}

	logger.Info("Delete user successfully", zap.String("userId", userId), zap.String("journey", "DeleteUser"))

	c.Status(http.StatusOK)
}
