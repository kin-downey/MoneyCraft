package models

import (
	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	UserID uint
	User User `gorm:"foreignKey:UserID; references:ID"`
	Amount string `json:"amount" binding:"required"`
}
