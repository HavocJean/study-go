package main

import (
	"context"
	"log"

	"github.com/HavocJean/study-go/internal/config/database/mongodb"
	"github.com/HavocJean/study-go/internal/logger"
	"github.com/HavocJean/study-go/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("Start use APP")

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error = %s \n", err.Error())
		return
	}

	userController := initDependenies(database)

	router := gin.Default()
	routes.InitiRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8082"); err != nil {
		logger.Error("Fail to start use port", err)
	}
}
