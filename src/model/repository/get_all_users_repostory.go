package repository

import (
	"context"
	"fmt"
	"os"

	"example.com/m/src/configuration/logger"
	"example.com/m/src/configuration/rest_err"
	"example.com/m/src/model"
	"example.com/m/src/model/repository/entity"
	"example.com/m/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) GetAllUsers() ([]model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init GetAllUsers repository", zap.String("journey", "GetAllUsers"))

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	var users []*entity.UserEntity

	object, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintln("No users found")
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		logger.Error("Error trying to find users", err, zap.String("journey", "GetAllUsers"))
		errorMessage := "Error trying to find users"
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	for object.Next(context.Background()) {
		userEntity := &entity.UserEntity{}
		err := object.Decode(userEntity)
		if err != nil {
			logger.Error("Error trying to decode user", err, zap.String("journey", "GetAllUsers"))
			errorMessage := "Error trying to decode user"
			return nil, rest_err.NewInternalServerError(errorMessage)
		}
		users = append(users, userEntity)
	}

	logger.Info("End GetAllUsers repository", zap.String("journey", "GetAllUsers"))

	var domains []model.UserDomainInterface
	for _, user := range users {
		domains = append(domains, converter.ConvertEntityToDomain(*user))
	}

	return domains, nil
}
