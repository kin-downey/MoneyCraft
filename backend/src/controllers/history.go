package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/y-watagashi/MoneyCraft/models"
)

type HistoryHandler struct {
	Db *gorm.DB
}

// Historyの新規作成と保存
func (h *HistoryHandler) AddHistory(c *gin.Context) {
    // ユーザーの検索
	var user models.User
	err := h.Db.First(&user, c.Param("id")).Error; if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User not found"})
		return
	}

	// 新しいHistoryレコードの作成
	history := models.History{
		UserID: user.ID,
		Amount: user.Currency,
	}

	// Historyレコードの保存
	if err := h.Db.Create(&history).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error saving history"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "history": history})
}