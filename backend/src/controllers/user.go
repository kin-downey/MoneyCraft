package controllers

import (	
	// "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/y-watagashi/MoneyCraft/models"
)

type UserHandler struct {
	Db *gorm.DB
}

func (h *UserHandler) AddUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user); if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	err = h.Db.Create(&user).Error; if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}