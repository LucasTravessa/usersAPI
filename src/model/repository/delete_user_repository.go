package repository

import (
	"context"
	"os"

	"example.com/m/src/configuration/logger"
	"example.com/m/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(
	userId string,
) *rest_err.RestErr {
	logger.Info("Init DeleteUser repository", zap.String("journey", "DeleteUser"))

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error while deleting  user", err, zap.String("journey", "DeleteUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("End DeleteUser repository", zap.String("journey", "DeleteUser"))

	return nil

}
