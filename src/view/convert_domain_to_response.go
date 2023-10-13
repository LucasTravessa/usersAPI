package view

import (
	"example.com/m/src/controller/model/response"
	"example.com/m/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID: "",
		Email: userDomain.GetEmail(),
		Name: userDomain.GetName(),
		Age: userDomain.GetAge(),
	}
}
