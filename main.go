package main

import (
	"github.com/HavocJean/study-go/internal/config/database/mongodb"
	"github.com/HavocJean/study-go/internal/controller"
	"github.com/HavocJean/study-go/internal/logger"
	"github.com/HavocJean/study-go/internal/model/service"
	"github.com/HavocJean/study-go/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("Start use APP")

	mongodb.InitiConnectMongodb()

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitiRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8082"); err != nil {
		logger.Error("Fail to start use port", err)
	}
}
