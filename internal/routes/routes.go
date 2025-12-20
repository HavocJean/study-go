package routes

import (
	"github.com/HavocJean/study-go/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitiRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/users/:userId", userController.FindUserByID)
	r.GET("/users/email/:email", userController.FindUserByEmail)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users", userController.UpdateUser)
	r.DELETE("/delete", userController.DeleteUser)
}
