package main

import (
	"log"

	"github.com/HavocJean/study-go/internal/controller"
	"github.com/HavocJean/study-go/internal/logger"
	"github.com/HavocJean/study-go/internal/model/service"
	"github.com/HavocJean/study-go/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Start use APP")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitiRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8082"); err != nil {
		logger.Error("Fail to start use port", err)
	}
}
