package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/HavocJean/study-go/internal/config/rest_error"
)

func NewUserDomain(email, password, name string, age int8) *userDomain {
	return &userDomain{
		email, password, name, age,
	}
}

type userDomain struct {
	// id       string
	Email    string
	Password string
	Name     string
	Age      int8
}

func (u *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_error.RestError
	UpdateUser(string) *rest_error.RestError
	FindUser(string) *rest_error.RestError
	DeleteUser(string) *rest_error.RestError
}
