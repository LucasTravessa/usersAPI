package model

import (
	"fmt"
	"os"
	"strings"
	"time"

	"example.com/m/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"age":   ud.age,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_err.NewInternalServerError(
			fmt.Sprintf("Error when generating token, err = %s", err.Error()))
	}

	return tokenString, nil
}

func VerifyToken(tokenValue string) (UserDomainInterface, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, rest_err.NewBadRequestError("Invalid token")
	})
	if err != nil {
		return nil, rest_err.NewUnauthorizedError("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, rest_err.NewUnauthorizedError("Invalid token")
	}

	ud := &userDomain{
		id:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   claims["age"].(int8),
	}

	return ud, nil
}

func VerifyTokenMiddleware(c *gin.Context) {
	secret := os.Getenv(JWT_SECRET_KEY)
	tokenValue := c.GetHeader("Authorization")

	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, rest_err.NewBadRequestError("Invalid token")
	})
	if err != nil {
		errRest := rest_err.NewUnauthorizedError("Invalid token")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		errRest := rest_err.NewUnauthorizedError("Invalid token")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

}

func RemoveBearerPrefix(token string) string {
	token = strings.TrimPrefix(token, "Bearer ")
	return token
}
