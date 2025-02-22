package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Login       string
	Password    string
	ValidClicks uint
	AllClicks   uint
}
