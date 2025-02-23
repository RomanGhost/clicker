package model

import "gorm.io/gorm"

type Transfer struct {
	gorm.Model
	SenderID   uint `gorm:"not null"` // Внешний ключ на User
	Sender     User `gorm:"foreignKey:SenderID"`
	ReceiverID uint `gorm:"not null"` // Внешний ключ на User
	Receiver   User `gorm:"foreignKey:ReceiverID"`
	Valid      uint `gorm:"not null"`
	Clicks     uint `gorm:"not null"`
}
