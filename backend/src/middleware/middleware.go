package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/y-watagashi/MoneyCraft/utils"
)

func AuthMiddleware(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "No token provided",
		})
		c.Abort()
		return
	}
	_, err := utils.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		c.Abort()
		return
	}
	c.Next()
}
