package database

import (
	"chat-back/database/model"
	"sync"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

// Функция получения экземпляра БД (Singleton)
func GetDBInstance(filepath string) *gorm.DB {
	once.Do(func() {
		var err error
		dbInstance, err = gorm.Open(sqlite.Open(filepath), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		// Автоматическая миграция моделей
		dbInstance.AutoMigrate(
			&model.User{},
			&model.League{},
			&model.Transaction{},
			&model.TranslationLeague{},
			&model.Language{},
		)

		// Запуск инициализации данных
		InitializeDatabase(dbInstance)
	})

	return dbInstance
}

// Главная функция инициализации данных
func InitializeDatabase(db *gorm.DB) {
	initializeLanguages(db)
	initializeLeagues(db)
	initializeTranslations(db)
}

func initializeLanguages(db *gorm.DB) {
	languages := []model.Language{
		{Language: "Русский"},
		{Language: "English"},
	}

	for _, language := range languages {
		db.FirstOrCreate(&language, model.Language{Language: language.Language})
	}
}

func initializeLeagues(db *gorm.DB) {
	leagues := []model.League{
		{Code: "Pizza Delivery", MinUsualClicks: 0},
		{Code: "Delivery", MinUsualClicks: 10_000},
		{Code: "Student", MinUsualClicks: 100_000},
		{Code: "Scientist", MinUsualClicks: 1_000_000},
		{Code: "Banker", MinUsualClicks: 10_000_000},
		{Code: "Investor", MinUsualClicks: 100_000_000},
		{Code: "Elite", MinUsualClicks: 1_000_000_000},
		{Code: "MLord", MinUsualClicks: 10_000_000_000},
	}

	for _, league := range leagues {
		db.FirstOrCreate(&league, model.League{Code: league.Code})
	}
}

func initializeTranslations(db *gorm.DB) {
	translations := []model.TranslationLeague{
		// Pizza Delivery
		{Name: "Доставщик", Description: "Начни с малого!", LanguageID: "Русский", LeagueID: 1},
		{Name: "Pizza Delivery", Description: "Start small!", LanguageID: "English", LeagueID: 1},

		// Delivery
		{Name: "Межпланетная перевозка", Description: "Теперь ты можешь доставлять больше!", LanguageID: "Русский", LeagueID: 2},
		{Name: "Planet Express", Description: "Now you can deliver more!", LanguageID: "English", LeagueID: 2},

		// Student
		{Name: "Студент", Description: "Пора учиться!", LanguageID: "Русский", LeagueID: 3},
		{Name: "Student", Description: "Time to study!", LanguageID: "English", LeagueID: 3},

		// Scientist
		{Name: "Ученый", Description: "Исследуй новые горизонты!", LanguageID: "Русский", LeagueID: 4},
		{Name: "Scientist", Description: "Explore new horizons!", LanguageID: "English", LeagueID: 4},

		// Banker
		{Name: "Банкир", Description: "Финансовый успех близок!", LanguageID: "Русский", LeagueID: 5},
		{Name: "Banker", Description: "Financial success is near!", LanguageID: "English", LeagueID: 5},

		// Investor
		{Name: "Инвестор", Description: "Вкладывай и приумножай!", LanguageID: "Русский", LeagueID: 6},
		{Name: "Investor", Description: "Invest and multiply!", LanguageID: "English", LeagueID: 6},

		// Elite
		{Name: "Элита", Description: "Ты достиг вершины!", LanguageID: "Русский", LeagueID: 7},
		{Name: "Elite", Description: "You have reached the top!", LanguageID: "English", LeagueID: 7},

		// MLord
		{Name: "Мем Лорд", Description: "Ты покорил интернет!", LanguageID: "Русский", LeagueID: 8},
		{Name: "Meme Lord", Description: "You've conquered the internet!", LanguageID: "English", LeagueID: 8},
	}

	for _, translation := range translations {
		db.FirstOrCreate(&translation, model.TranslationLeague{Name: translation.Name, LanguageID: translation.LanguageID, LeagueID: translation.LeagueID})
	}
}
