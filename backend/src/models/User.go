package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	IsParent bool `json:"isParent" binding:"required"`
	currency int `json:"currency" binding:"required"`
	FamilyID uint
}

