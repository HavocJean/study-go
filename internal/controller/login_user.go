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

func (u *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init loginUser controller", zap.String("journey", "loginUser"))

	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error binding login user controller", err, zap.String("journey", "loginUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(int(errRest.Code), errRest)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)

	domainResult, err := u.service.LoginUserServices(domain)
	if err != nil {
		logger.Error("Error in login user controller", err, zap.String("journey", "loginUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(int(errRest.Code), errRest)
		return
	}

	logger.Info("Finish loginUser controller", zap.String("journey", "loginUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))
}
