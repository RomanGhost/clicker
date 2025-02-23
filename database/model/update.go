package model

type Update struct {
	ID              uint   `gorm:"primaryKey"`
	Name            string `gorm:"unique;not null"`
	MaxLevel        uint
	ValidCoef       float64
	ClickCoef       float64
	MinLeague       uint
	Price           uint
	PriceGrowthCoef float32
}
