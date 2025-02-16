package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Login       string
	ValidClicks uint
	AllClicks   uint
}
