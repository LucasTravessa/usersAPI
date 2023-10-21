package service

import (
	"example.com/m/src/configuration/rest_err"
	"example.com/m/src/model"
	"example.com/m/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	GetAllUsersService() ([]model.UserDomainInterface, *rest_err.RestErr)
	CreateUserService(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
	FindUserByIDServices(id string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr)
	LoginUserService(model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr)
}
