package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/y-watagashi/MoneyCraft/utils"
	"github.com/y-watagashi/MoneyCraft/models"
)

type UserHandler struct {
	Db *gorm.DB
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var user models.User
	fmt.Println(c.Request.Body)
	err := c.ShouldBindJSON(&user); if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "error": err.Error()})
		return
	}
	// 家族IDを追加
	user.FamilyID = 1
	// JWTトークンを生成
	tokenString, err := utils.GenerateToken(user.ID); if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to generate token"})
		return
	}
	// ユーザーを作成
	err = h.Db.Create(&user).Error; if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "access_token": tokenString, "user": user})
}


func (h *UserHandler) GetUser(c *gin.Context) {
	var user models.User
	err := h.Db.First(&user, c.Param("id")).Error; if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "user": user})
}