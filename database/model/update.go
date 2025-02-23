package model

type Update struct {
	ID        uint    `gorm:"primaryKey"`
	Name      string  `gorm:"unique;not null"`
	MaxLevel  uint    `gorm:"not null"`
	ValidCoef float64 `gorm:"not null"`
	ClickCoef float64 `gorm:"not null"`
}
