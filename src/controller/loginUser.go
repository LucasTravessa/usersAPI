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

func (uc *userControllerInterface) LoginUser(c *gin.Context) {

	logger.Info("Init LoginUser", zap.String("journey", "LoginUser"))

	var userRequest request.UserLogin
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "LoginUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)

	domainResult, token, err := uc.service.LoginUserService(domain)
	if err != nil {
		logger.Error("Error trying to call login service", err, zap.String("journey", "LoginUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully", zap.String("journey", "LoginUser"), zap.Any("user", view.ConvertDomainToResponse(domain)))

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
