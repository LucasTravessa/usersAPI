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
	GetAllUsers() ([]model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailAndPassword(email string, password string) (model.UserDomainInterface, *rest_err.RestErr)
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(userId string) *rest_err.RestErr
}
