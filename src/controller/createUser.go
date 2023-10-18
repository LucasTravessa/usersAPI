package controller

import (
	"net/http"

	"example.com/m/src/configuration/logger"
	"example.com/m/src/configuration/validation"
	"example.com/m/src/controller/model/request"
	"example.com/m/src/model"
	"example.com/m/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {

	logger.Info("Init CreateUser", zap.String("journey", "CreateUser"))

	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "CreateUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	domainResult, err := uc.service.CreateUserService(domain)
	if err != nil {
		logger.Error("Error trying to create user", err, zap.String("journey", "CreateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully", zap.String("journey", "CreateUser"), zap.Any("user", view.ConvertDomainToResponse(domain)))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
