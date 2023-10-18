package main

import (
	"context"
	"log"

	"example.com/m/src/configuration/database/mongodb"
	"example.com/m/src/controller"
	"example.com/m/src/controller/routes"
	"example.com/m/src/model/repository"
	"example.com/m/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
		return
	}

	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	controller := controller.NewUserControllerInterface(service)

	// userController := initDependencies(database)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, controller)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
