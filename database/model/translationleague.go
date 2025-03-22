package model

type TranslationLeague struct {
	ID          int      `gorm:"primaryKey"`
	Name        string   `gorm:"not null;unique"`
	Description string   `gorm:"not null"`
	League      League   `gorm:"foreignKey:LeagueID"`
	LeagueID    int      `gorm:"default:null"`
	Language    Language `gorm:"foreignKey:LanguageID"`
	LanguageID  string   `gorm:"size:32"`
}
