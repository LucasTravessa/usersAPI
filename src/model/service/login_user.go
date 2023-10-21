package service

import (
	"example.com/m/src/configuration/logger"
	"example.com/m/src/configuration/rest_err"
	"example.com/m/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info("Init LoginUser service", zap.String("journey", "LoginUser"))

	userDomain.EncryptPassword()

	user, err := ud.findUserByEmailAndPasswordServices(userDomain.GetEmail(), userDomain.GetPassword())

	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	logger.Info("Successfully created user", zap.String("journey", "LoginUser"))
	return user, token, nil
}
