package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marcofilho/go-crud-api/src/controller"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/getUserByID/:userID", userController.FindUserByID)
	r.GET("/getUserByEmail/:email", userController.FindUserByEmail)
	r.GET("getAllUsers/", userController.FindAllUsers)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userID", userController.UpdateUser)
	r.DELETE("/deleteUser/:userID", userController.DeleteUser)
}
