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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmailAndPassword(email string, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmailAndPassword repository", zap.String("journey", "findUserByEmailAndPassword"))

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "User not found with this email and password"
			return nil, rest_err.NewForbiddenError(errorMessage)
		}
		logger.Error("Error trying to find user by email and password", err, zap.String("journey", "findUserByEmailAndPassword"))
		errorMessage := "Error trying to find user by email and password"
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("End findUserByEmailAndPassword repository", zap.String("journey", "findUserByEmailAndPassword"), zap.String("email", userEntity.Email))

	return converter.ConvertEntityToDomain(*userEntity), nil

}

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmil repository", zap.String("journey", "findUserByEmail"))

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		logger.Error("Error trying to find user by email", err, zap.String("journey", "findUserByEmail"))
		errorMessage := "Error trying to find user by email"
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("End findUserByEmil repository", zap.String("journey", "findUserByEmail"), zap.String("email", userEntity.Email))

	return converter.ConvertEntityToDomain(*userEntity), nil

}

func (ur *userRepository) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserById repository", zap.String("journey", "findUserByID"))

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	obectId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: obectId}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this ID: %s", id)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		logger.Error("Error trying to find user by ID", err, zap.String("journey", "findUserByID"))
		errorMessage := "Error trying to find user by ID"
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("End findUserByID repository", zap.String("journey", "findUserByID"), zap.String("userID", userEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(*userEntity), nil

}
