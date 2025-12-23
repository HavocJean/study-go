package service

import (
	"github.com/HavocJean/study-go/internal/config/rest_error"
	"github.com/HavocJean/study-go/internal/logger"
	"github.com/HavocJean/study-go/internal/model"
	"go.uber.org/zap"
)

func (u *userDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_error.RestError) {
	logger.Info("Init FindUserByEmailServices service", zap.String("journey", "FindUserByEmailServices"))

	return u.userRepository.FindUserByEmail(email)
}

func (u *userDomainService) FindUserByIDServices(id string) (model.UserDomainInterface, *rest_error.RestError) {
	logger.Info("Init FindUserByIdServices service", zap.String("journey", "FindUserByIdServices"))

	return u.userRepository.FindUserByID(id)
}

func (u *userDomainService) findUserByEmailAndPasswordServices(email, password string) (model.UserDomainInterface, *rest_error.RestError) {
	logger.Info("Init findUserByEmailAndPasswordServices service", zap.String("journey", "findUserByEmailAndPasswordServices"))

	return u.userRepository.FindUserByEmailAndPassword(email, password)
}
