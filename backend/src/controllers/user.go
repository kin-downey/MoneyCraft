package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/y-watagashi/MoneyCraft/utils"
	"github.com/y-watagashi/MoneyCraft/models"
	"golang.org/x/crypto/bcrypt"
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
	user.Currency = 0
	// JWTトークンを生成
	tokenString, err := utils.GenerateToken(user.ID); if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to generate token"})
		return
	}
	// ユーザー作成の際にパスワードをハッシュ化する
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10); if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)
	// ユーザーをDBに保存
	err = h.Db.Create(&user).Error; if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "access_token": tokenString, "user": user})
}


func (h *UserHandler) SignIn(c *gin.Context) {
	var reqBody struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
	// リクエストボディからユーザー情報を取得
	err := c.ShouldBindJSON(&reqBody); if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "error": err.Error()})
		return
	}
	// ユーザーをDBから取得
	var user models.User
	err = h.Db.Where("email = ?", reqBody.Email).First(&user).Error; if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get user"})
		return
	}
	// パスワードをチェック
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password)); if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to compare password"})
		return
	}
	// JWTトークンを生成
	tokenString, err := utils.GenerateToken(user.ID); if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to generate token"})
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