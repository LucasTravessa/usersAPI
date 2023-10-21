package routes

import (
	"example.com/m/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface,
) {
	r.GET("/users", userController.GetAllUsers)
	r.GET("/userById/:userId", userController.FindUserById)
	r.GET("/userByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PATCH("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)

	r.POST("/login", userController.LoginUser)

}
