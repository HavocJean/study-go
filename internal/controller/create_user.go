package controller

import (
	"net/http"

	"github.com/HavocJean/study-go/internal/config/validation"
	"github.com/HavocJean/study-go/internal/controller/model/request"
	"github.com/HavocJean/study-go/internal/logger"
	"github.com/HavocJean/study-go/internal/model"
	"github.com/HavocJean/study-go/internal/model/service"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
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

	service := service.NewUserDomainService()
	if err := service.CreateUser(domain); err != nil {
		c.JSON(int(err.Code), err)
		return
	}

	logger.Info("User created successfully")

	c.String(http.StatusOK, "")
}
