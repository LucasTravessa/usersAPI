package repository

import (
	"example.com/m/src/configuration/rest_err"
	"example.com/m/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		db,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
}
