package main

import (
	"log"

	"github.com/HavocJean/study-go/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	routes.InitiRoutes(&router.RouterGroup)

	if err := router.Run(":8082"); err != nil {
		log.Fatal(err)
	}
}
