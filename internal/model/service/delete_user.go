package service

import (
	"github.com/HavocJean/study-go/internal/config/rest_error"
	"github.com/HavocJean/study-go/internal/logger"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUserServices(userId string) *rest_error.RestError {
	logger.Info("Init DeleteUserService service", zap.String("journey", "deleteUserService"))

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "deleteUserService"))
		return err
	}

	logger.Info(
		"DeleteUserService service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUserService"),
	)

	return nil
}
