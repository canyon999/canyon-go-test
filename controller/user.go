package controller

import (
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	ResponseSuccess(c, gin.H{
		"accessToken": "aToken",
	})
}
