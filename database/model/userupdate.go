package model

import "gorm.io/gorm"

type UserUpdate struct {
	gorm.Model
	UserID    uint   `gorm:"primaryKey"`
	UpdateID  uint   `gorm:"primaryKey"`
	User      User   `gorm:"foreignKey:UserID"`
	Update    Update `gorm:"foreignKey:UpdateID"`
	Level     uint
	ClickCoef float64 `gorm:"default:1"`
	ValidCoef float64 `gorm:"default:1"`
}
