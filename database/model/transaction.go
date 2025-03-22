package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	SenderID   int `gorm:"not null"`
	ReceiverID int `gorm:"not null"`
	Valid      float64
	Clicks     float64
	Sender     User `gorm:"foreignKey:SenderID"`
	Receiver   User `gorm:"foreignKey:ReceiverID"`
}
