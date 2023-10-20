package service

import (
	"example.com/m/src/configuration/logger"
	"example.com/m/src/configuration/rest_err"
	"example.com/m/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init UpdateUser service", zap.String("journey", "UpdateUser"))
	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error while creating user", err, zap.String("journey", "UpdateUser"))

		return err
	}
	logger.Info("Successfully created user", zap.String("journey", "UpdateUser"))
	return nil
}
