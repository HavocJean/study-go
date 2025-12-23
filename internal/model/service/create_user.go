package service

import (
	"github.com/HavocJean/study-go/internal/config/rest_error"
	"github.com/HavocJean/study-go/internal/logger"
	"github.com/HavocJean/study-go/internal/model"
	"go.uber.org/zap"
)

func (u *userDomainService) CreateUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_error.RestError) {
	logger.Info("Init CreateUser service", zap.String("journey", "createUser"))

	user, _ := u.FindUserByEmailServices(userDomain.GetEmail())
	if user != nil {
		return nil, rest_error.NewBadRequestError("Email is already registered")
	}

	userDomain.EncryptPassword()

	userDomainRepository, err := u.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info(
		"CreateUser service executed successfully",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("journey", "createUser"),
	)

	return userDomainRepository, nil
}
