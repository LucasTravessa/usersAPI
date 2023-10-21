package controller

import (
	"net/http"

	"example.com/m/src/configuration/logger"
	"example.com/m/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser", zap.String("journey", "DeleteUser"))

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid user id")
		c.JSON(errRest.Code, errRest)
	}

	err := uc.service.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to delete user", err, zap.String("journey", "DeleteUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User deleted successfully", zap.String("journey", "DeleteUser"), zap.String("userId", userId))
	c.Status(http.StatusOK)

}
