package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required" gorm:"unique;not null;"`
	Password string `json:"password" binding:"required"`
	IsParent bool 
	Currency int 
	FamilyID uint
}