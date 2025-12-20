package controller

import (
	"net/http"
	"net/mail"

	"github.com/HavocJean/study-go/internal/config/rest_error"
	"github.com/HavocJean/study-go/internal/logger"
	"github.com/HavocJean/study-go/internal/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init FindUserByEmail controller", zap.String("journey", "FindUserByEmail"))

	email := c.Param("email")

	if _, err := mail.ParseAddress(email); err != nil {
		errorMessage := rest_error.NewBadRequestError("email is not a valid email")
		logger.Error("Invalid user email format", err, zap.String("email", email))
		c.JSON(int(errorMessage.Code), errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(email)
	if err != nil {
		logger.Error("Error try findUserByEmail service", err, zap.String("email", email))
		c.JSON(int(err.Code), err)
		return
	}

	logger.Info("FindUserByEmail controller successfully", zap.String("email", "FindUserByEmail"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))

}

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init FindUserByID controller", zap.String("journey", "FindUserByID"))

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMessage := rest_error.NewBadRequestError("userId is not a valid id")
		logger.Error("Invalid user ID format", err, zap.String("userId", userId))
		c.JSON(int(errorMessage.Code), errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error try findUserById service", err, zap.String("userId", userId))
		c.JSON(int(err.Code), err)
		return
	}

	logger.Info("FindUserByID controller successfully", zap.String("userId", "FindUserByID"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
