package controller

import (
	"net/http"

	"github.com/HavocJean/study-go/internal/config/rest_error"
	"github.com/HavocJean/study-go/internal/config/validation"
	"github.com/HavocJean/study-go/internal/controller/model/request"
	"github.com/HavocJean/study-go/internal/logger"
	"github.com/HavocJean/study-go/internal/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser controller", zap.String("journey", "updateUser"))

	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "updateUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(int(errRest.Code), errRest)
		return
	}

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_error.NewBadRequestError("Invalid userid, must be a hex value")
		c.JSON(int(errRest.Code), errRest)
		return
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateUserServices(userId, domain)
	if err != nil {
		logger.Error("Error try to call UpdateUser service", err, zap.String("journey", "updateUser"))

		c.JSON(int(err.Code), err)
		return
	}

	logger.Info("Update user successfully", zap.String("userId", userId), zap.String("journey", "updateUser"))

	c.Status(http.StatusOK)
}
