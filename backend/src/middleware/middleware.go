package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/y-watagashi/MoneyCraft/utils"
)

func AuthMiddleware(c *gin.Context) {
	tokenString, err := c.Request.Header.Get("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}
	token, err := utils.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		c.Abort()
		return
	}
	c.Next()
}
