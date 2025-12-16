package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string

	SetID(string)

	GetJSONValue() (string, error)

	EncryptPassword()
}

func NewUserDomain(email, password, name string, age int8) *userDomain {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

func (ud *userDomain) SetID(id string) {
	ud.ID = id
}

type userDomain struct {
	ID       string
	Email    string `json:"email"`
	Password string
	Name     string
	Age      int8
}

func (ud *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(b), nil
}

func (u *userDomain) GetEmail() string {
	return u.Email
}

func (u *userDomain) GetPassword() string {
	return u.Password
}

func (u *userDomain) GetName() string {
	return u.Name
}

func (u *userDomain) GetAge() int8 {
	return u.Age
}

func (u *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hash.Sum(nil))
}
