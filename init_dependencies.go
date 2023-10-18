package main

import (
	"example.com/m/src/controller"
	"example.com/m/src/model/repository"
	"example.com/m/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserContrllerInterface(service)

}
