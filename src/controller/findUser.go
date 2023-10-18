package controller

import (
	"net/http"
	"net/mail"

	"example.com/m/src/configuration/logger"
	"example.com/m/src/configuration/rest_err"
	"example.com/m/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init findUserByID Controller", zap.String("journey", "findUserByID"))

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userId", err, zap.String("journey", "findUserByID"))
		errorMessage := rest_err.NewBadRequestError("invalid user id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error trying to call findUserByID services", err, zap.String("journey", "findUserByID"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Finish findUserByIDServices Controller", zap.String("journey", "findUserByID"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))

}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail Controller", zap.String("journey", "findUserByEmail"))

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validate userEmail", err, zap.String("journey", "findUserByEmail"))
		errorMessage := rest_err.NewBadRequestError("invalid user email")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to call findUserByEmail services", err, zap.String("journey", "findUserByEmail"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Finish findUserByEmailServices Controller", zap.String("journey", "findUserByEmail"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
