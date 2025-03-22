package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Login       string   `gorm:"unique;not null"`
	Password    string   `gorm:"not null"`
	ValidClicks float64  `gorm:"default:0"`
	UsualClicks float64  `gorm:"default:0"`
	League      League   `gorm:"foreignKey:LeagueID"`
	LeagueID    int      `gorm:"default:1"`
	Language    Language `gorm:"foreignKey:LanguageID"`
	LanguageID  string   `gorm:"size:32;default:English"`
}
