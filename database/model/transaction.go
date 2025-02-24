package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	SenderID   uint `gorm:"not null"` // Внешний ключ на User
	Sender     User `gorm:"foreignKey:SenderID"`
	ReceiverID uint `gorm:"not null"` // Внешний ключ на User
	Receiver   User `gorm:"foreignKey:ReceiverID"`
	Valid      float64
	Clicks     float64
}
