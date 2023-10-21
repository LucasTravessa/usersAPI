package service

import (
	"example.com/m/src/configuration/logger"
	"example.com/m/src/configuration/rest_err"
	"example.com/m/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByIDServices service", zap.String("journey", "findUserByIDServices"))

	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmailServices service", zap.String("journey", "findUserByEmailServices"))

	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) findUserByEmailAndPasswordServices(email string, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmailAndPasswordServices service", zap.String("journey", "findUserByEmailAndPasswordServices"))

	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
