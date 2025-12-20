package converter

import (
	"github.com/HavocJean/study-go/internal/model"
	"github.com/HavocJean/study-go/internal/model/repository/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	objID, _ := primitive.ObjectIDFromHex(domain.GetID())

	return &entity.UserEntity{
		ID:       objID,
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}
