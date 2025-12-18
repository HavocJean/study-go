package controller

import (
	"net/http"

	"github.com/HavocJean/study-go/internal/config/validation"
	"github.com/HavocJean/study-go/internal/controller/model/request"
	"github.com/HavocJean/study-go/internal/logger"
	"github.com/HavocJean/study-go/internal/model"
	"github.com/HavocJean/study-go/internal/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(int(errRest.Code), errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	domainResult, err := uc.service.CreateUserServices(domain)
	if err != nil {
		logger.Error("Error try to call CreateUser service", err, zap.String("journey", "createUser"))

		c.JSON(int(err.Code), err)
		return
	}

	logger.Info("User created successfully", zap.String("userId", domainResult.GetID()))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
