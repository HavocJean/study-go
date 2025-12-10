package service

import (
	"fmt"

	"github.com/HavocJean/study-go/internal/config/rest_error"
	"github.com/HavocJean/study-go/internal/logger"
	"github.com/HavocJean/study-go/internal/model"
	"go.uber.org/zap"
)

func (u *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_error.RestError {
	logger.Info("Init CreateUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}
