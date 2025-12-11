package controller

import (
	"net/http"

	"github.com/HavocJean/study-go/internal/config/validation"
	"github.com/HavocJean/study-go/internal/controller/model/request"
	"github.com/HavocJean/study-go/internal/logger"
	"github.com/HavocJean/study-go/internal/model"
	"github.com/HavocJean/study-go/internal/view"
	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
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

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(int(err.Code), err)
		return
	}

	logger.Info("User created successfully")

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}
