package models

import (
	"gorm.io/gorm"
)

type Family struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	User []User
}