package controller

import "github.com/gin-gonic/gin"

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	c.String(200, "FindUserById")
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	c.String(200, "FindUserByEmail")
}
