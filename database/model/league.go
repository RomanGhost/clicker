package model

type League struct {
	Number uint   `gorm:"primaryKey"`
	Name   string `gorm:"unique;not null"`
}
