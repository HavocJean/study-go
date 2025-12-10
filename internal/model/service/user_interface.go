package service

import (
	"github.com/HavocJean/study-go/internal/config/rest_error"
	"github.com/HavocJean/study-go/internal/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_error.RestError
	UpdateUser(string, model.UserDomainInterface) *rest_error.RestError
	FindUser(string) (*model.UserDomainInterface, *rest_error.RestError)
	DeleteUser(string) *rest_error.RestError
}
