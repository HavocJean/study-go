package service

import (
	"github.com/HavocJean/study-go/internal/config/rest_error"
	"github.com/HavocJean/study-go/internal/logger"
	"github.com/HavocJean/study-go/internal/model"
	"go.uber.org/zap"
)

func (u *userDomainService) UpdateUserServices(userId string, userDomain model.UserDomainInterface) *rest_error.RestError {
	logger.Info("Init UpdateUserService service", zap.String("journey", "updateUserService"))

	err := u.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "updateUserService"))
		return err
	}

	logger.Info(
		"UpdateUserService service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUserService"),
	)

	return nil
}
