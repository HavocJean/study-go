package model

import (
	"fmt"

	"github.com/HavocJean/study-go/internal/config/rest_error"
	"github.com/HavocJean/study-go/internal/logger"
	"go.uber.org/zap"
)

var (
	UserInterface UserDomainInterface
)

func (u *userDomain) CreateUser() *rest_error.RestError {
	logger.Info("Init CreateUser model", zap.String("journey", "createUser"))

	u.EncryptPassword()

	fmt.Println(u)

	return nil
}
