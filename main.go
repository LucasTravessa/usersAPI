package main

import (
	"log"

	"example.com/m/src/controller"
	"example.com/m/src/controller/routes"
	"example.com/m/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserContrllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
