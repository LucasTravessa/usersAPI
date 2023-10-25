package routes

import (
	"example.com/m/src/controller"
	"example.com/m/src/model"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface,
) {
	r.GET("/users", model.VerifyTokenMiddleware, userController.GetAllUsers)
	r.GET("/userById/:userId", model.VerifyTokenMiddleware, userController.FindUserById)
	r.GET("/userByEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", model.VerifyTokenMiddleware, userController.CreateUser)
	r.PATCH("/updateUser/:userId", model.VerifyTokenMiddleware, userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", model.VerifyTokenMiddleware, userController.DeleteUser)

	r.POST("/login", userController.LoginUser)

}
