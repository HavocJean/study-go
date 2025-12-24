package service

import (
	"github.com/HavocJean/study-go/internal/config/rest_error"
	"github.com/HavocJean/study-go/internal/logger"
	"github.com/HavocJean/study-go/internal/model"
	"go.uber.org/zap"
)

func (u *userDomainService) LoginUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_error.RestError) {
	logger.Info("Init LoginUser service", zap.String("journey", "LoginUser"))

	userDomain.EncryptPassword()

	user, err := u.findUserByEmailAndPasswordServices(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	logger.Info(
		"LoginUser service executed successfully",
		zap.String("userId", user.GetID()),
		zap.String("journey", "LoginUser"),
	)

	return user, token, nil
}
