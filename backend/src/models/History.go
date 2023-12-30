package models

import (
	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	UserID uint
	Amount string `json:"amount" binding:"required"`
}
