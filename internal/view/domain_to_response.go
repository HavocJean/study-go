package view

import (
	"github.com/HavocJean/study-go/internal/controller/model/response"
	"github.com/HavocJean/study-go/internal/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    nil,
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
