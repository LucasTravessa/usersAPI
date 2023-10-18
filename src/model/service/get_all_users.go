package service

import (
	"example.com/m/src/configuration/logger"
	"example.com/m/src/configuration/rest_err"
	"example.com/m/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) GetAllUsersService() ([]model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init GetAllUsersServices service", zap.String("journey", "GetAllUsersService"))

	return ud.userRepository.GetAllUsers()
}
