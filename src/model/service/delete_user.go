package service

import (
	"example.com/m/src/configuration/logger"
	"example.com/m/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("Init DeleteUser service", zap.String("journey", "DeleteUser"))
	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error while creating user", err, zap.String("journey", "DeleteUser"))

		return err
	}
	logger.Info("Successfully created user", zap.String("journey", "DeleteUser"))
	return nil
}
