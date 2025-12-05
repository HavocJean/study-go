package routes

import (
	"github.com/HavocJean/study-go/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitiRoutes(r *gin.RouterGroup) {

	r.GET("/users/:id", controller.GetUserById)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users", controller.UpdateUser)
	r.DELETE("/delete", controller.DeleteUser)
}
