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

func GetDBInstance(filepath string) *gorm.DB {
	once.Do(func() {
		var err error
		dbInstance, err = gorm.Open(sqlite.Open(filepath), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		// Автоматическая миграция модели
		dbInstance.AutoMigrate(
			&model.User{},
			&model.League{},
			&model.Update{},
			&model.UserUpdate{},
			&model.Transfer{},
		)
	})

	addLeagueInfo(dbInstance)
	addUpdatesInfo(dbInstance)
	return dbInstance
}

func addLeagueInfo(db *gorm.DB) {
	leagues := []model.League{
		{Number: 1, Name: "Новобранец OWCA"},
		{Number: 2, Name: "Агент-стажёр"},
		{Number: 3, Name: "Оперативник"},
		{Number: 4, Name: "Фитнес-Эксель"},
		{Number: 5, Name: "Супер-агент OWCA"},
		{Number: 6, Name: "Элитный Оперативник"},
		{Number: 7, Name: "Гранд-Мастер Агент"},
		{Number: 8, Name: "Легенда OWCA"},
	}

	for _, league := range leagues {
		db.FirstOrCreate(&league, model.League{Number: league.Number})
	}
}

func addUpdatesInfo(db *gorm.DB) {
	updates := []model.Update{
		// Улучшения для лиги 1 (MinLeague = 1, бонус на валидные клики отсутствует)
		{Name: "Эспандер Пэрри", MaxLevel: 10, ValidCoef: 1.0, ClickCoef: 1.05, MinLeague: 1, Price: 100, PriceGrowthCoef: 1.10},
		{Name: "Костюм шпиона", MaxLevel: 5, ValidCoef: 1.0, ClickCoef: 1.04, MinLeague: 1, Price: 100, PriceGrowthCoef: 1.10},
		{Name: "Фитнес-браслет OWCA", MaxLevel: 5, ValidCoef: 1.0, ClickCoef: 1.05, MinLeague: 1, Price: 110, PriceGrowthCoef: 1.10},
		{Name: "Умные часы агента", MaxLevel: 4, ValidCoef: 1.0, ClickCoef: 1.03, MinLeague: 1, Price: 110, PriceGrowthCoef: 1.10},
		{Name: "Спортивный костюм Pro", MaxLevel: 6, ValidCoef: 1.0, ClickCoef: 1.06, MinLeague: 1, Price: 120, PriceGrowthCoef: 1.10},
		{Name: "Биосканер активности", MaxLevel: 3, ValidCoef: 1.0, ClickCoef: 1.07, MinLeague: 1, Price: 120, PriceGrowthCoef: 1.10},
		{Name: "Динамический ускоритель", MaxLevel: 6, ValidCoef: 1.0, ClickCoef: 1.08, MinLeague: 1, Price: 130, PriceGrowthCoef: 1.10},
		{Name: "Виртуальный тренер", MaxLevel: 5, ValidCoef: 1.0, ClickCoef: 1.05, MinLeague: 1, Price: 130, PriceGrowthCoef: 1.10},
		{Name: "Нейро-активатор", MaxLevel: 4, ValidCoef: 1.0, ClickCoef: 1.04, MinLeague: 1, Price: 140, PriceGrowthCoef: 1.10},
		{Name: "Протокол быстроты", MaxLevel: 3, ValidCoef: 1.0, ClickCoef: 1.05, MinLeague: 1, Price: 140, PriceGrowthCoef: 1.10},

		// Улучшения для лиги 2 (MinLeague = 2)
		{Name: "Ботинки с пружинами", MaxLevel: 7, ValidCoef: 1.0, ClickCoef: 1.07, MinLeague: 2, Price: 150, PriceGrowthCoef: 1.12},
		{Name: "Гантели-энергия", MaxLevel: 8, ValidCoef: 1.0, ClickCoef: 1.08, MinLeague: 2, Price: 150, PriceGrowthCoef: 1.12},
		{Name: "Рюкзак быстродействия", MaxLevel: 5, ValidCoef: 1.0, ClickCoef: 1.07, MinLeague: 2, Price: 160, PriceGrowthCoef: 1.12},
		{Name: "Пульсовый модификатор", MaxLevel: 4, ValidCoef: 1.0, ClickCoef: 1.09, MinLeague: 2, Price: 160, PriceGrowthCoef: 1.12},
		{Name: "Молниеносный стимулятор", MaxLevel: 5, ValidCoef: 1.0, ClickCoef: 1.10, MinLeague: 2, Price: 170, PriceGrowthCoef: 1.12},
		{Name: "Сила через кроссовки", MaxLevel: 4, ValidCoef: 1.0, ClickCoef: 1.08, MinLeague: 2, Price: 170, PriceGrowthCoef: 1.12},
		{Name: "Электро-энергия", MaxLevel: 5, ValidCoef: 1.0, ClickCoef: 1.09, MinLeague: 2, Price: 180, PriceGrowthCoef: 1.12},
		{Name: "Турбо-кроссовки", MaxLevel: 3, ValidCoef: 1.0, ClickCoef: 1.10, MinLeague: 2, Price: 180, PriceGrowthCoef: 1.12},
		{Name: "Стартовый импульс", MaxLevel: 4, ValidCoef: 1.0, ClickCoef: 1.11, MinLeague: 2, Price: 190, PriceGrowthCoef: 1.12},
		{Name: "Двойной клик", MaxLevel: 2, ValidCoef: 1.0, ClickCoef: 1.12, MinLeague: 2, Price: 190, PriceGrowthCoef: 1.12},

		// Улучшения для лиги 3 (MinLeague = 3)
		{Name: "Шпионские очки OWCA", MaxLevel: 3, ValidCoef: 1.0, ClickCoef: 1.10, MinLeague: 3, Price: 200, PriceGrowthCoef: 1.14},
		{Name: "Бандана скрытности", MaxLevel: 3, ValidCoef: 1.0, ClickCoef: 1.08, MinLeague: 3, Price: 200, PriceGrowthCoef: 1.14},
		{Name: "Массажный валик", MaxLevel: 5, ValidCoef: 1.0, ClickCoef: 1.09, MinLeague: 3, Price: 210, PriceGrowthCoef: 1.14},
		{Name: "Спортивный дрон", MaxLevel: 5, ValidCoef: 1.0, ClickCoef: 1.10, MinLeague: 3, Price: 210, PriceGrowthCoef: 1.14},
		{Name: "Прокачка нейронов", MaxLevel: 4, ValidCoef: 1.0, ClickCoef: 1.11, MinLeague: 3, Price: 220, PriceGrowthCoef: 1.14},
		{Name: "Голографический тренер", MaxLevel: 3, ValidCoef: 1.0, ClickCoef: 1.12, MinLeague: 3, Price: 220, PriceGrowthCoef: 1.14},
		{Name: "Нейро-оптимизатор", MaxLevel: 4, ValidCoef: 1.0, ClickCoef: 1.11, MinLeague: 3, Price: 230, PriceGrowthCoef: 1.14},
		{Name: "Бустер интеллекта", MaxLevel: 3, ValidCoef: 1.0, ClickCoef: 1.12, MinLeague: 3, Price: 230, PriceGrowthCoef: 1.14},
		{Name: "Шоковый регулятор", MaxLevel: 2, ValidCoef: 1.0, ClickCoef: 1.13, MinLeague: 3, Price: 240, PriceGrowthCoef: 1.14},
		{Name: "Режим активации", MaxLevel: 4, ValidCoef: 1.0, ClickCoef: 1.14, MinLeague: 3, Price: 240, PriceGrowthCoef: 1.14},

		// Улучшения для лиги 4 (MinLeague = 4)
		{Name: "Реактивный ранец Монограмма", MaxLevel: 4, ValidCoef: 1.0, ClickCoef: 1.12, MinLeague: 4, Price: 250, PriceGrowthCoef: 1.16},
		{Name: "Пояс силы", MaxLevel: 6, ValidCoef: 1.0, ClickCoef: 1.13, MinLeague: 4, Price: 250, PriceGrowthCoef: 1.16},
		{Name: "Фитбол тренировки", MaxLevel: 4, ValidCoef: 1.0, ClickCoef: 1.12, MinLeague: 4, Price: 260, PriceGrowthCoef: 1.16},
		{Name: "Силовая перчатка", MaxLevel: 7, ValidCoef: 1.0, ClickCoef: 1.14, MinLeague: 4, Price: 260, PriceGrowthCoef: 1.16},
		{Name: "Суперфит шлем", MaxLevel: 3, ValidCoef: 1.0, ClickCoef: 1.15, MinLeague: 4, Price: 270, PriceGrowthCoef: 1.16},
		{Name: "Генератор мощности", MaxLevel: 5, ValidCoef: 1.0, ClickCoef: 1.16, MinLeague: 4, Price: 270, PriceGrowthCoef: 1.16},
		{Name: "Тактический ремень", MaxLevel: 4, ValidCoef: 1.0, ClickCoef: 1.15, MinLeague: 4, Price: 280, PriceGrowthCoef: 1.16},
		{Name: "Энергия в динамике", MaxLevel: 3, ValidCoef: 1.0, ClickCoef: 1.17, MinLeague: 4, Price: 280, PriceGrowthCoef: 1.16},
		{Name: "Импульс тренировки", MaxLevel: 4, ValidCoef: 1.0, ClickCoef: 1.18, MinLeague: 4, Price: 290, PriceGrowthCoef: 1.16},
		{Name: "Стратегический бустер", MaxLevel: 3, ValidCoef: 1.0, ClickCoef: 1.19, MinLeague: 4, Price: 290, PriceGrowthCoef: 1.16},

		// Улучшения для лиги 5 (MinLeague = 5)
		{Name: "Мультигаджет от Карла", MaxLevel: 2, ValidCoef: 1.0, ClickCoef: 1.20, MinLeague: 5, Price: 300, PriceGrowthCoef: 1.18},
		{Name: "Бутылка энергии", MaxLevel: 5, ValidCoef: 1.0, ClickCoef: 1.21, MinLeague: 5, Price: 300, PriceGrowthCoef: 1.18},
		{Name: "Турник суперсилы", MaxLevel: 6, ValidCoef: 1.0, ClickCoef: 1.22, MinLeague: 5, Price: 310, PriceGrowthCoef: 1.18},
		{Name: "Быстрые ласты", MaxLevel: 4, ValidCoef: 1.0, ClickCoef: 1.23, MinLeague: 5, Price: 310, PriceGrowthCoef: 1.18},
		{Name: "Электронный бодибилдер", MaxLevel: 5, ValidCoef: 1.0, ClickCoef: 1.24, MinLeague: 5, Price: 320, PriceGrowthCoef: 1.18},
		{Name: "Фитнес-генератор", MaxLevel: 4, ValidCoef: 1.0, ClickCoef: 1.25, MinLeague: 5, Price: 320, PriceGrowthCoef: 1.18},
		{Name: "Турбосила", MaxLevel: 3, ValidCoef: 1.0, ClickCoef: 1.26, MinLeague: 5, Price: 330, PriceGrowthCoef: 1.18},
		{Name: "Взрывной заряд", MaxLevel: 2, ValidCoef: 1.0, ClickCoef: 1.27, MinLeague: 5, Price: 330, PriceGrowthCoef: 1.18},
		{Name: "Сила активности", MaxLevel: 4, ValidCoef: 1.0, ClickCoef: 1.28, MinLeague: 5, Price: 340, PriceGrowthCoef: 1.18},
		{Name: "Импульс к эффективности", MaxLevel: 3, ValidCoef: 1.0, ClickCoef: 1.29, MinLeague: 5, Price: 340, PriceGrowthCoef: 1.18},

		// Улучшения для лиги 6 (MinLeague = 6, теперь бонус на валидные клики начинает действовать)
		{Name: "Интерактивный тренер", MaxLevel: 5, ValidCoef: 1.10, ClickCoef: 1.30, MinLeague: 6, Price: 350, PriceGrowthCoef: 1.20},
		{Name: "Фитнес-плита", MaxLevel: 3, ValidCoef: 1.10, ClickCoef: 1.31, MinLeague: 6, Price: 350, PriceGrowthCoef: 1.20},
		{Name: "Энергетический пульсометр", MaxLevel: 4, ValidCoef: 1.10, ClickCoef: 1.32, MinLeague: 6, Price: 360, PriceGrowthCoef: 1.20},
		{Name: "Энергетический импульс", MaxLevel: 3, ValidCoef: 1.10, ClickCoef: 1.33, MinLeague: 6, Price: 360, PriceGrowthCoef: 1.20},
		{Name: "Стимулятор активности", MaxLevel: 4, ValidCoef: 1.10, ClickCoef: 1.34, MinLeague: 6, Price: 370, PriceGrowthCoef: 1.20},
		{Name: "Бустер выносливости", MaxLevel: 3, ValidCoef: 1.10, ClickCoef: 1.35, MinLeague: 6, Price: 370, PriceGrowthCoef: 1.20},
		{Name: "Калорийный ускоритель", MaxLevel: 2, ValidCoef: 1.10, ClickCoef: 1.36, MinLeague: 6, Price: 380, PriceGrowthCoef: 1.20},
		{Name: "Тренировочный эксель", MaxLevel: 4, ValidCoef: 1.10, ClickCoef: 1.37, MinLeague: 6, Price: 380, PriceGrowthCoef: 1.20},
		{Name: "Фитнес-бустер", MaxLevel: 3, ValidCoef: 1.10, ClickCoef: 1.38, MinLeague: 6, Price: 390, PriceGrowthCoef: 1.20},
		{Name: "Пульсовый импульс", MaxLevel: 2, ValidCoef: 1.10, ClickCoef: 1.39, MinLeague: 6, Price: 390, PriceGrowthCoef: 1.20},

		// Улучшения для лиги 7 (MinLeague = 7)
		{Name: "Секретный протокол OWCA", MaxLevel: 3, ValidCoef: 1.20, ClickCoef: 1.40, MinLeague: 7, Price: 400, PriceGrowthCoef: 1.25},
		{Name: "Зарядный коврик", MaxLevel: 5, ValidCoef: 1.20, ClickCoef: 1.41, MinLeague: 7, Price: 400, PriceGrowthCoef: 1.25},
		{Name: "Антигравитационные шорты", MaxLevel: 4, ValidCoef: 1.20, ClickCoef: 1.42, MinLeague: 7, Price: 410, PriceGrowthCoef: 1.25},
		{Name: "Голографический имбустер", MaxLevel: 3, ValidCoef: 1.20, ClickCoef: 1.43, MinLeague: 7, Price: 410, PriceGrowthCoef: 1.25},
		{Name: "Кибернетический наручник", MaxLevel: 5, ValidCoef: 1.20, ClickCoef: 1.44, MinLeague: 7, Price: 420, PriceGrowthCoef: 1.25},
		{Name: "Магнитный кардиомонитор", MaxLevel: 3, ValidCoef: 1.20, ClickCoef: 1.45, MinLeague: 7, Price: 420, PriceGrowthCoef: 1.25},
		{Name: "Рефлекторный комбинезон", MaxLevel: 4, ValidCoef: 1.20, ClickCoef: 1.46, MinLeague: 7, Price: 430, PriceGrowthCoef: 1.25},
		{Name: "Активатор сенсоров", MaxLevel: 2, ValidCoef: 1.20, ClickCoef: 1.47, MinLeague: 7, Price: 430, PriceGrowthCoef: 1.25},
		{Name: "Силовой имбустер", MaxLevel: 3, ValidCoef: 1.20, ClickCoef: 1.48, MinLeague: 7, Price: 440, PriceGrowthCoef: 1.25},
		{Name: "Прокачка адреналина", MaxLevel: 2, ValidCoef: 1.20, ClickCoef: 1.49, MinLeague: 7, Price: 440, PriceGrowthCoef: 1.25},

		// Улучшения для лиги 8 (MinLeague = 8)
		{Name: "Агентский арсенал", MaxLevel: 1, ValidCoef: 1.30, ClickCoef: 1.50, MinLeague: 8, Price: 500, PriceGrowthCoef: 1.30},
		{Name: "Легендарный Эликсир", MaxLevel: 1, ValidCoef: 1.30, ClickCoef: 1.51, MinLeague: 8, Price: 500, PriceGrowthCoef: 1.30},
		{Name: "Турбощит", MaxLevel: 6, ValidCoef: 1.30, ClickCoef: 1.52, MinLeague: 8, Price: 510, PriceGrowthCoef: 1.30},
		{Name: "Ускоритель сердцебиения", MaxLevel: 5, ValidCoef: 1.30, ClickCoef: 1.53, MinLeague: 8, Price: 510, PriceGrowthCoef: 1.30},
		{Name: "Оптический тренажер", MaxLevel: 4, ValidCoef: 1.30, ClickCoef: 1.54, MinLeague: 8, Price: 520, PriceGrowthCoef: 1.30},
		{Name: "Импульс квантум", MaxLevel: 3, ValidCoef: 1.30, ClickCoef: 1.55, MinLeague: 8, Price: 520, PriceGrowthCoef: 1.30},
		{Name: "Кибернетический имплант", MaxLevel: 4, ValidCoef: 1.30, ClickCoef: 1.56, MinLeague: 8, Price: 530, PriceGrowthCoef: 1.30},
		{Name: "Экзоскелет агента", MaxLevel: 5, ValidCoef: 1.30, ClickCoef: 1.57, MinLeague: 8, Price: 530, PriceGrowthCoef: 1.30},
		{Name: "Сверхзвуковой протокол", MaxLevel: 2, ValidCoef: 1.30, ClickCoef: 1.58, MinLeague: 8, Price: 540, PriceGrowthCoef: 1.30},
		{Name: "Финальный драйв", MaxLevel: 1, ValidCoef: 1.30, ClickCoef: 1.59, MinLeague: 8, Price: 540, PriceGrowthCoef: 1.30},
	}

	for _, up := range updates {
		db.FirstOrCreate(&up, model.Update{Name: up.Name})
	}
}
