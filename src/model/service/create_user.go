package service

import (
	"example.com/m/src/configuration/logger"
	"example.com/m/src/configuration/rest_err"
	"example.com/m/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()
	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error while creating user", err, zap.String("journey", "createUser"))

		return nil, err
	}
	logger.Info("Successfully created user", zap.String("journey", "createUser"))
	return userDomainRepository, nil
}
