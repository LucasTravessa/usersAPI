package controller

import "github.com/gin-gonic/gin"

func FindUserById(c *gin.Context) {
	c.String(200, "FindUserById")
}

func FindUserByEmail(c *gin.Context) {
	c.String(200, "FindUserByEmail")
}