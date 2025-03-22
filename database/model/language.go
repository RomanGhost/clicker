package model

type Language struct {
	Language string `gorm:"primaryKey;size:32;not null"`
}
