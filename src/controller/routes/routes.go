package routes

import (
	"example.com/m/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup){
	r.GET("/userById/:userId", controller.FindUserById)
	r.GET("/userByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdadeUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
