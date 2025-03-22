package model

type League struct {
	ID             uint   `gorm:"primaryKey"`
	Code           string `gorm:"unique;not null"`
	MinUsualClicks int
	MinValidClicks int `gorm:"dafault:0"`
}
