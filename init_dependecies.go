package main

import (
	"github.com/HavocJean/study-go/internal/controller"
	"github.com/HavocJean/study-go/internal/model/repository"
	"github.com/HavocJean/study-go/internal/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependenies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
