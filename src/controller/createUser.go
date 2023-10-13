package controller

import (
	"net/http"

	"example.com/m/src/configuration/logger"
	"example.com/m/src/configuration/validation"
	"example.com/m/src/controller/model/request"
	"example.com/m/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {

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

	if err := domain.CreateUser(); err != nil {
		logger.Error("Error trying to create user", err, zap.String("journey", "CreateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully", zap.String("journey", "CreateUser"))
	c.JSON(http.StatusOK, "")
}
