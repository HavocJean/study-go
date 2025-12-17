package service

import (
	"github.com/HavocJean/study-go/internal/config/rest_error"
	"github.com/HavocJean/study-go/internal/model"
	"github.com/HavocJean/study-go/internal/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_error.RestError

	UpdateUser(string, model.UserDomainInterface) *rest_error.RestError
	FindUser(string) (*model.UserDomainInterface, *rest_error.RestError)
	DeleteUser(string) *rest_error.RestError
}
