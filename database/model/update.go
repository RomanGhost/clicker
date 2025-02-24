package model

type Update struct {
	ID              uint   `gorm:"primaryKey"`
	Name            string `gorm:"unique;not null"`
	MaxLevel        uint
	ValidCoef       float64
	ClickCoef       float64
	MinLeague       uint
	PriceGrowthCoef float32
	PriceValid      float64
	PriceClick      float64
	Description     string
}
