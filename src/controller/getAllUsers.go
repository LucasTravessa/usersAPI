package controller

import (
	"fmt"
	"net/http"

	"example.com/m/src/configuration/logger"
	"example.com/m/src/model"
	"example.com/m/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) GetAllUsers(c *gin.Context) {
	logger.Info("Init getAllUsers Controller", zap.String("journey", "getAllUsers"))

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))

	if err != nil {
		logger.Error("Error trying to verify token", err, zap.String("journey", "getAllUsers"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("User authenticated: %v", user))

	userDomain, err := uc.service.GetAllUsersService()
	if err != nil {
		logger.Error("Error trying to call getAllUsers services", err, zap.String("journey", "getAllUsers"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Finish getAllUsers Controller", zap.String("journey", "getAllUsers"))

	c.JSON(http.StatusOK, view.ConvertDomainsToResponses(userDomain))

}
