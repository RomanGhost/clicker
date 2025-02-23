package model

import "gorm.io/gorm"

type League struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}
