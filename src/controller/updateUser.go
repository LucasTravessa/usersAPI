package controller

import (
	"net/http"

	"example.com/m/src/configuration/logger"
	"example.com/m/src/configuration/rest_err"
	"example.com/m/src/configuration/validation"
	"example.com/m/src/controller/model/request"
	"example.com/m/src/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {

	logger.Info("Init UpdateUser", zap.String("journey", "UpdateUser"))

	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "UpdateUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid user id")
		c.JSON(errRest.Code, errRest)
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Error("Error trying to update user", err, zap.String("journey", "UpdateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User updated successfully", zap.String("journey", "UpdateUser"), zap.String("userId", userId))
	c.Status(http.StatusOK)

}
