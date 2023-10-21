package service

import (
	"example.com/m/src/configuration/logger"
	"example.com/m/src/configuration/rest_err"
	"example.com/m/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser service", zap.String("journey", "createUser"))

	user, _ := ud.FindUserByEmailServices(userDomain.GetEmail())

	if user != nil {
		return nil, rest_err.NewBadRequestError("Email already exist")
	}

	userDomain.EncryptPassword()
	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error while creating user", err, zap.String("journey", "createUser"))

		return nil, err
	}
	logger.Info("Successfully created user", zap.String("journey", "createUser"))
	return userDomainRepository, nil
}
