package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	FromUserID uint
	ToUserID uint
	FromUser User `gorm:"foreignKey:FromUserID; references:ID"`
	ToUser User `gorm:"foreignKey:ToUserID; references:ID"`
	Amount int
}