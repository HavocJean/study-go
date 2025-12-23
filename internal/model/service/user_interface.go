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
	CreateUserServices(model.UserDomainInterface) (model.UserDomainInterface, *rest_error.RestError)

	FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_error.RestError)
	FindUserByIDServices(id string) (model.UserDomainInterface, *rest_error.RestError)

	UpdateUserServices(string, model.UserDomainInterface) *rest_error.RestError
	DeleteUserServices(string) *rest_error.RestError

	LoginUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_error.RestError)
}
