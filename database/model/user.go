package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Login       string  `gorm:"unique;not null"`
	Password    string  `gorm:"not null"`
	ValidClicks float64 `gorm:"default:0"`
	AllClicks   float64 `gorm:"default:0"`
	LeagueID    uint    // Внешний ключ для связи с League
	League      League  `gorm:"foreignKey:LeagueID"`
}
