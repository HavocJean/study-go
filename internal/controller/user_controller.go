package controller

import (
	"github.com/HavocJean/study-go/internal/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserByEmail(c *gin.Context)
	FindUserByID(c *gin.Context)

	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)

	LoginUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
