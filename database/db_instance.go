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
			&model.Transaction{},
		)
	})

	addLeagueInfo(dbInstance)
	addUpdatesInfo(dbInstance)
	return dbInstance
}

func addLeagueInfo(db *gorm.DB) {
	leagues := []model.League{
		{Number: 1, Name: "Новобранец OWCA", MinClicks: 0},
		{Number: 2, Name: "Агент-стажёр", MinClicks: 10_000},
		{Number: 3, Name: "Оперативник", MinClicks: 1_000_000},
		{Number: 4, Name: "Фитнес-Эксель", MinClicks: 10_000_000},
		{Number: 5, Name: "Супер-агент OWCA", MinClicks: 1_000_000_000},
		{Number: 6, Name: "Элитный Оперативник", MinClicks: 5_000_000_000},
		{Number: 7, Name: "Гранд-Мастер Агент", MinClicks: 10_000_000_000},
		{Number: 8, Name: "Легенда OWCA", MinClicks: 100_000_000_000},
	}

	for _, league := range leagues {
		db.FirstOrCreate(&league, model.League{Number: league.Number})
	}
}

func addUpdatesInfo(db *gorm.DB) {
	updates := []model.Update{
		// Лига 1 (Новобранец OWCA) - Базовое снаряжение для начинающих агентов
		{
			Name: "Маскировка 'Хамелеон'", MaxLevel: 999, ValidCoef: 1.0, ClickCoef: 1.02, MinLeague: 1,
			PriceGrowthCoef: 1.07, PriceClick: 500, PriceValid: 0,
			Description: "Простейший камуфляж, чтобы слиться с толпой.",
		},
		{
			Name: "Шпионский бинокль 2x", MaxLevel: 999, ValidCoef: 1.0, ClickCoef: 1.03, MinLeague: 1,
			PriceGrowthCoef: 1.08, PriceClick: 800, PriceValid: 0,
			Description: "Базовый бинокль для наблюдения за объектами.",
		},
		{
			Name: "Курс 'Быстрый набор'", MaxLevel: 999, ValidCoef: 1.0, ClickCoef: 1.04, MinLeague: 1,
			PriceGrowthCoef: 1.09, PriceClick: 1200, PriceValid: 0,
			Description: "Ускоренный курс подготовки для начинающих агентов.",
		},
		{
			Name: "Агентский блокнот", MaxLevel: 999, ValidCoef: 1.0, ClickCoef: 1.05, MinLeague: 1,
			PriceGrowthCoef: 1.10, PriceClick: 1800, PriceValid: 0,
			Description: "Блокнот с ручкой для важных заметок и планов.",
		},
		{
			Name: "Стартовый паек агента", MaxLevel: 5, ValidCoef: 1.0, ClickCoef: 1.2, MinLeague: 1,
			PriceGrowthCoef: 1.12, PriceClick: 7500, PriceValid: 3,
			Description: "Необходимый запас провизии для выполнения первых заданий.",
		},

		// Лига 2 (Агент-стажёр) - Специализация и улучшенное снаряжение
		{
			Name: "Ботинки 'Скорость тени'", MaxLevel: 1999, ValidCoef: 1.0, ClickCoef: 1.06, MinLeague: 2,
			PriceGrowthCoef: 1.10, PriceClick: 8000, PriceValid: 0,
			Description: "Легкие ботинки для быстрого и бесшумного передвижения.",
		},
		{
			Name: "Усиленные гантели", MaxLevel: 1999, ValidCoef: 1.0, ClickCoef: 1.07, MinLeague: 2,
			PriceGrowthCoef: 1.11, PriceClick: 12000, PriceValid: 0,
			Description: "Тяжелые гантели для поддержания физической формы.",
		},
		{
			Name: "Рюкзак 'Неуловимость'", MaxLevel: 1999, ValidCoef: 1.0, ClickCoef: 1.08, MinLeague: 2,
			PriceGrowthCoef: 1.12, PriceClick: 18000, PriceValid: 0,
			Description: "Вместительный рюкзак с функциями маскировки.",
		},
		{
			Name: "Модуль 'Пульс'", MaxLevel: 1999, ValidCoef: 1.0, ClickCoef: 1.09, MinLeague: 2,
			PriceGrowthCoef: 1.13, PriceClick: 25000, PriceValid: 0,
			Description: "Устройство для контроля пульса и уровня стресса.",
		},
		{
			Name: "Стимулятор 'Молния'", MaxLevel: 1999, ValidCoef: 1.0, ClickCoef: 1.10, MinLeague: 2,
			PriceGrowthCoef: 1.14, PriceClick: 35000, PriceValid: 0,
			Description: "Инъектор для кратковременного увеличения скорости реакции.",
		},
		{
			Name: "Кроссовки 'Сила воли'", MaxLevel: 999, ValidCoef: 1.0, ClickCoef: 1.12, MinLeague: 2,
			PriceGrowthCoef: 1.15, PriceClick: 50000, PriceValid: 1,
			Description: "Специальные кроссовки, повышающие выносливость.",
		},
		{
			Name: "Энергетический гель 'Вольт'", MaxLevel: 999, ValidCoef: 1.0, ClickCoef: 1.13, MinLeague: 2,
			PriceGrowthCoef: 1.16, PriceClick: 70000, PriceValid: 1,
			Description: "Гель для быстрого восстановления энергии.",
		},
		{
			Name: "Турбо-кроссовки 'Форсаж'", MaxLevel: 999, ValidCoef: 1.0, ClickCoef: 1.14, MinLeague: 2,
			PriceGrowthCoef: 1.17, PriceClick: 95000, PriceValid: 1,
			Description: "Кроссовки с турбо-ускорением для рывков.",
		},
		{
			Name: "Импульсный пистолет 'Старт'", MaxLevel: 999, ValidCoef: 1.0, ClickCoef: 1.15, MinLeague: 2,
			PriceGrowthCoef: 1.18, PriceClick: 120000, PriceValid: 5,
			Description: "Пистолет, генерирующий импульс для ускорения действий.",
		},
		{
			Name: "Программа 'Двойной клик'", MaxLevel: 5, ValidCoef: 1.0, ClickCoef: 1.20, MinLeague: 2,
			PriceGrowthCoef: 1.20, PriceClick: 900_000, PriceValid: 250,
			Description: "Улучшенная программа тренировок для удвоения эффективности кликов.",
		},

		// Лига 3 (Оперативник) - Высокотехнологичные гаджеты и нейросети
		{
			Name: "Шпионские очки OWCA", MaxLevel: 5999, ValidCoef: 1.0, ClickCoef: 1.12, MinLeague: 3,
			PriceGrowthCoef: 1.12, PriceClick: 900_000, PriceValid: 5,
			Description: "Очки с рентгеновским и ночным видением.",
		},
		{
			Name: "Бандана скрытности", MaxLevel: 5999, ValidCoef: 1.0, ClickCoef: 1.13, MinLeague: 3,
			PriceGrowthCoef: 1.13, PriceClick: 500_000, PriceValid: 5,
			Description: "Бандана, создающая поле невидимости малой дальности.",
		},
		{
			Name: "Массажный валик", MaxLevel: 5999, ValidCoef: 1.0, ClickCoef: 1.14, MinLeague: 3,
			PriceGrowthCoef: 1.14, PriceClick: 650_000, PriceValid: 3,
			Description: "Массажер для быстрого снятия мышечного напряжения.",
		},
		{
			Name: "Спортивный дрон", MaxLevel: 5999, ValidCoef: 1.0, ClickCoef: 1.15, MinLeague: 3,
			PriceGrowthCoef: 1.15, PriceClick: 800_000, PriceValid: 8,
			Description: "Миниатюрный дрон для разведки и поддержки.",
		},
		{
			Name: "Прокачка нейронов", MaxLevel: 1000, ValidCoef: 1.0, ClickCoef: 1.16, MinLeague: 3,
			PriceGrowthCoef: 1.05, PriceClick: 1_000_000, PriceValid: 2,
			Description: "Нейросеть, оптимизирующая мозговую деятельность.", // Оставил MaxLevel 1000 как было, особенное улучшение
		},
		{
			Name: "Голографический тренер", MaxLevel: 5999, ValidCoef: 1.0, ClickCoef: 1.17, MinLeague: 3,
			PriceGrowthCoef: 1.16, PriceClick: 1_200_000, PriceValid: 3,
			Description: "Голографический тренер для индивидуальных тренировок.",
		},
		{
			Name: "Нейро-оптимизатор", MaxLevel: 5999, ValidCoef: 1.0, ClickCoef: 1.18, MinLeague: 3,
			PriceGrowthCoef: 1.17, PriceClick: 1_500_000, PriceValid: 10,
			Description: "Устройство для нейронной оптимизации и ускорения обучения.",
		},
		{
			Name: "Бустер интеллекта", MaxLevel: 5999, ValidCoef: 1.0, ClickCoef: 1.19, MinLeague: 3,
			PriceGrowthCoef: 1.18, PriceClick: 1_800_000, PriceValid: 25,
			Description: "Бустер для кратковременного повышения интеллектуальных способностей.",
		},
		{
			Name: "Шоковый регулятор", MaxLevel: 5999, ValidCoef: 1.0, ClickCoef: 1.20, MinLeague: 3,
			PriceGrowthCoef: 1.19, PriceClick: 2_200_000, PriceValid: 60,
			Description: "Устройство для управления уровнем адреналина и стресса.",
		},
		{
			Name: "Режим активации", MaxLevel: 5, ValidCoef: 1.0, ClickCoef: 1.25, MinLeague: 3,
			PriceGrowthCoef: 1.25, PriceClick: 9_000_000, PriceValid: 90,
			Description: "Эксклюзивный протокол активации максимальной производительности.", // Оставил MaxLevel 5 как было, особенное улучшение
		},

		// Лига 4 (Фитнес-Эксель) - Экзоскелеты и продвинутые системы
		{
			Name: "Реактивный ранец Монограмма", MaxLevel: 899, ValidCoef: 1.0, ClickCoef: 1.20, MinLeague: 4,
			PriceGrowthCoef: 1.14, PriceClick: 8_000_000, PriceValid: 100,
			Description: "Легкий реактивный ранец для быстрого перемещения.",
		},
		{
			Name: "Пояс силы Титан", MaxLevel: 899, ValidCoef: 1.0, ClickCoef: 1.21, MinLeague: 4,
			PriceGrowthCoef: 1.14, PriceClick: 8_500_000, PriceValid: 105,
			Description: "Пояс, увеличивающий физическую силу и выносливость.",
		},
		{
			Name: "Фитбол тренировки", MaxLevel: 899, ValidCoef: 1.0, ClickCoef: 1.22, MinLeague: 4,
			PriceGrowthCoef: 1.14, PriceClick: 9_000_000, PriceValid: 105,
			Description: "Фитбол для тренировки баланса и координации.",
		},
		{
			Name: "Силовая перчатка Удар", MaxLevel: 899, ValidCoef: 1.0, ClickCoef: 1.23, MinLeague: 4,
			PriceGrowthCoef: 1.14, PriceClick: 9_500_000, PriceValid: 110,
			Description: "Перчатка, усиливающая силу удара в несколько раз.",
		},
		{
			Name: "Суперфит шлем Фокус", MaxLevel: 5, ValidCoef: 1.0, ClickCoef: 1.24, MinLeague: 4,
			PriceGrowthCoef: 1.14, PriceClick: 10_000_000, PriceValid: 110,
			Description: "Шлем, улучшающий концентрацию и реакцию.",
		},

		// Лига 5 (Супер-агент OWCA) - Био-импланты и энергетические системы
		{
			Name: "Мультигаджет от Карла", MaxLevel: 2999, ValidCoef: 1.0, ClickCoef: 1.25, MinLeague: 5,
			PriceGrowthCoef: 1.16, PriceClick: 50_000_000, PriceValid: 120,
			Description: "Универсальный гаджет с набором инструментов и функций.",
		},
		{
			Name: "Бутылка энергии", MaxLevel: 2999, ValidCoef: 1.0, ClickCoef: 1.26, MinLeague: 5,
			PriceGrowthCoef: 1.16, PriceClick: 55_000_000, PriceValid: 125,
			Description: "Напиток, обеспечивающий мощный прилив энергии.",
		},
		{
			Name: "Турник суперсилы", MaxLevel: 2999, ValidCoef: 1.0, ClickCoef: 1.27, MinLeague: 5,
			PriceGrowthCoef: 1.16, PriceClick: 60_000_000, PriceValid: 130,
			Description: "Турник, автоматически подстраивающий нагрузку.",
		},
		{
			Name: "Быстрые ласты", MaxLevel: 2999, ValidCoef: 1.0, ClickCoef: 1.28, MinLeague: 5,
			PriceGrowthCoef: 1.16, PriceClick: 65_000_000, PriceValid: 130,
			Description: "Ласты для невероятно быстрого плавания.",
		},
		{
			Name: "Электронный бодибилдер", MaxLevel: 2999, ValidCoef: 1.0, ClickCoef: 1.29, MinLeague: 5,
			PriceGrowthCoef: 1.16, PriceClick: 70_000_000, PriceValid: 135,
			Description: "Устройство для электростимуляции мышц.",
		},
		{
			Name: "Фитнес-генератор", MaxLevel: 2999, ValidCoef: 1.0, ClickCoef: 1.30, MinLeague: 5,
			PriceGrowthCoef: 1.16, PriceClick: 75_000_000, PriceValid: 135,
			Description: "Генератор, преобразующий движение в энергию.",
		},
		{
			Name: "Турбосила", MaxLevel: 2999, ValidCoef: 1.0, ClickCoef: 1.31, MinLeague: 5,
			PriceGrowthCoef: 1.16, PriceClick: 80_000_000, PriceValid: 140,
			Description: "Устройство для кратковременного увеличения всех физических параметров.",
		},
		{
			Name: "Взрывной заряд", MaxLevel: 2999, ValidCoef: 1.0, ClickCoef: 1.32, MinLeague: 5,
			PriceGrowthCoef: 1.16, PriceClick: 85_000_000, PriceValid: 140,
			Description: "Заряд энергии для максимального рывка.",
		},
		{
			Name: "Сила активности", MaxLevel: 2999, ValidCoef: 1.0, ClickCoef: 1.33, MinLeague: 5,
			PriceGrowthCoef: 1.16, PriceClick: 90_000_000, PriceValid: 145,
			Description: "Браслет, увеличивающий силу в зависимости от активности.",
		},
		{
			Name: "Импульс к эффективности", MaxLevel: 5, ValidCoef: 1.0, ClickCoef: 1.34, MinLeague: 5,
			PriceGrowthCoef: 1.16, PriceClick: 95_000_000, PriceValid: 145,
			Description: "Имплант, повышающий общую эффективность действий.",
		},

		// Лига 6 (Элитный Оперативник) - Продвинутые нейронные и энергетические системы, ValidCoef = 1.10
		{
			Name: "Интерактивный тренер", MaxLevel: 3999, ValidCoef: 1.10, ClickCoef: 1.35, MinLeague: 6,
			PriceGrowthCoef: 1.18, PriceClick: 200_000_000, PriceValid: 150,
			Description: "Тренер на основе нейросети, адаптирующийся к вашим потребностям.",
		},
		{
			Name: "Фитнес-плита", MaxLevel: 3999, ValidCoef: 1.10, ClickCoef: 1.36, MinLeague: 6,
			PriceGrowthCoef: 1.18, PriceClick: 210_000_000, PriceValid: 155,
			Description: "Платформа, изменяющая гравитацию для усиления тренировок.",
		},
		{
			Name: "Энергетический пульсометр", MaxLevel: 3999, ValidCoef: 1.10, ClickCoef: 1.37, MinLeague: 6,
			PriceGrowthCoef: 1.18, PriceClick: 220_000_000, PriceValid: 160,
			Description: "Пульсометр, анализирующий энергетические потоки организма.",
		},
		{
			Name: "Энергетический импульс", MaxLevel: 3999, ValidCoef: 1.10, ClickCoef: 1.38, MinLeague: 6,
			PriceGrowthCoef: 1.18, PriceClick: 230_000_000, PriceValid: 160,
			Description: "Мощный энергетический импульс, активирующий скрытые резервы.",
		},
		{
			Name: "Стимулятор активности", MaxLevel: 3999, ValidCoef: 1.10, ClickCoef: 1.39, MinLeague: 6,
			PriceGrowthCoef: 1.18, PriceClick: 240_000_000, PriceValid: 165,
			Description: "Стимулятор для поддержания максимальной активности в любых условиях.",
		},
		{
			Name: "Бустер выносливости", MaxLevel: 3999, ValidCoef: 1.10, ClickCoef: 1.40, MinLeague: 6,
			PriceGrowthCoef: 1.18, PriceClick: 250_000_000, PriceValid: 165,
			Description: "Бустер для экстремального увеличения выносливости.",
		},
		{
			Name: "Калорийный ускоритель", MaxLevel: 3999, ValidCoef: 1.10, ClickCoef: 1.41, MinLeague: 6,
			PriceGrowthCoef: 1.18, PriceClick: 260_000_000, PriceValid: 170,
			Description: "Устройство для ускоренного сжигания калорий и получения энергии.",
		},
		{
			Name: "Тренировочный эксель", MaxLevel: 3999, ValidCoef: 1.10, ClickCoef: 1.42, MinLeague: 6,
			PriceGrowthCoef: 1.18, PriceClick: 270_000_000, PriceValid: 170,
			Description: "Продвинутая программа тренировок для достижения пика формы.",
		},
		{
			Name: "Фитнес-бустер", MaxLevel: 3999, ValidCoef: 1.10, ClickCoef: 1.43, MinLeague: 6,
			PriceGrowthCoef: 1.18, PriceClick: 280_000_000, PriceValid: 175,
			Description: "Фитнес-бустер, выводящий физические возможности на уровень богов.",
		},
		{
			Name: "Пульсовый импульс", MaxLevel: 5, ValidCoef: 1.10, ClickCoef: 1.44, MinLeague: 6,
			PriceGrowthCoef: 1.18, PriceClick: 290_000_000, PriceValid: 175,
			Description: "Импульс, усиливающий сердцебиение и кровоток для максимальной производительности.",
		},

		// Лига 7 (Гранд-Мастер Агент) - Секретные разработки и уникальные протоколы, ValidCoef = 1.20
		{
			Name: "Секретный протокол OWCA", MaxLevel: 2999, ValidCoef: 1.20, ClickCoef: 1.45, MinLeague: 7,
			PriceGrowthCoef: 1.20, PriceClick: 700_000_000, PriceValid: 200,
			Description: "Секретный протокол для полного восстановления и регенерации.",
		},
		{
			Name: "Зарядный коврик", MaxLevel: 2999, ValidCoef: 1.20, ClickCoef: 1.46, MinLeague: 7,
			PriceGrowthCoef: 1.20, PriceClick: 720_000_000, PriceValid: 205,
			Description: "Коврик для беспроводной зарядки всех энергетических систем.",
		},
		{
			Name: "Антигравитационные шорты", MaxLevel: 2999, ValidCoef: 1.20, ClickCoef: 1.47, MinLeague: 7,
			PriceGrowthCoef: 1.20, PriceClick: 740_000_000, PriceValid: 210,
			Description: "Шорты, создающие антигравитационное поле для легкости движений.",
		},
		{
			Name: "Голографический имбустер", MaxLevel: 2999, ValidCoef: 1.20, ClickCoef: 1.48, MinLeague: 7,
			PriceGrowthCoef: 1.20, PriceClick: 760_000_000, PriceValid: 210,
			Description: "Имбустер, создающий голографические иллюзии для дезориентации противника.",
		},
		{
			Name: "Кибернетический наручник", MaxLevel: 2999, ValidCoef: 1.20, ClickCoef: 1.49, MinLeague: 7,
			PriceGrowthCoef: 1.20, PriceClick: 780_000_000, PriceValid: 215,
			Description: "Наруч для управления кибернетическими устройствами и нейросетями.",
		},
		{
			Name: "Магнитный кардиомонитор", MaxLevel: 2999, ValidCoef: 1.20, ClickCoef: 1.50, MinLeague: 7,
			PriceGrowthCoef: 1.20, PriceClick: 800_000_000, PriceValid: 215,
			Description: "Монитор, усиливающий и контролирующий сердцебиение на магнитном уровне.",
		},
		{
			Name: "Рефлекторный комбинезон", MaxLevel: 2999, ValidCoef: 1.20, ClickCoef: 1.51, MinLeague: 7,
			PriceGrowthCoef: 1.20, PriceClick: 820_000_000, PriceValid: 220,
			Description: "Комбинезон, усиливающий рефлексы до молниеносной скорости.",
		},
		{
			Name: "Активатор сенсоров", MaxLevel: 2999, ValidCoef: 1.20, ClickCoef: 1.52, MinLeague: 7,
			PriceGrowthCoef: 1.20, PriceClick: 840_000_000, PriceValid: 220,
			Description: "Устройство, усиливающее все органы чувств до предела.",
		},
		{
			Name: "Силовой имбустер", MaxLevel: 2999, ValidCoef: 1.20, ClickCoef: 1.53, MinLeague: 7,
			PriceGrowthCoef: 1.20, PriceClick: 860_000_000, PriceValid: 225,
			Description: "Имбустер, дающий силу легендарного титана Атланта.",
		},
		{
			Name: "Прокачка адреналина", MaxLevel: 5, ValidCoef: 1.20, ClickCoef: 1.54, MinLeague: 7,
			PriceGrowthCoef: 1.20, PriceClick: 880_000_000, PriceValid: 225,
			Description: "Система для мгновенной прокачки адреналина до критических уровней.",
		},

		// Лига 8 (Легенда OWCA) - Абсолютное превосходство и легендарное снаряжение, ValidCoef = 1.30
		{
			Name: "Арсенал Агент-Легенда", MaxLevel: 10, ValidCoef: 1.30, ClickCoef: 1.55, MinLeague: 8,
			PriceGrowthCoef: 1.25, PriceClick: 100_000_000_000, PriceValid: 250,
			Description: "Полный арсенал легендарного агента, включающий все необходимое.",
		},
		{
			Name: "Легендарный Эликсир", MaxLevel: 10, ValidCoef: 1.30, ClickCoef: 1.56, MinLeague: 8,
			PriceGrowthCoef: 1.25, PriceClick: 105_000_000_000, PriceValid: 255,
			Description: "Легендарный эликсир, дарующий невероятную живучесть.",
		},
		{
			Name: "Турбо-щит Непробиваемость", MaxLevel: 899, ValidCoef: 1.30, ClickCoef: 1.57, MinLeague: 8,
			PriceGrowthCoef: 1.25, PriceClick: 110_000_000_000, PriceValid: 260,
			Description: "Щит, способный выдержать любые атаки.",
		},
		{
			Name: "Ускоритель сердцебиения Пульсар", MaxLevel: 899, ValidCoef: 1.30, ClickCoef: 1.58, MinLeague: 8,
			PriceGrowthCoef: 1.25, PriceClick: 115_000_000_000, PriceValid: 260,
			Description: "Устройство, разгоняющее сердцебиение до космических скоростей.",
		},
		{
			Name: "Оптический тренажер Всевидение", MaxLevel: 899, ValidCoef: 1.30, ClickCoef: 1.59, MinLeague: 8,
			PriceGrowthCoef: 1.25, PriceClick: 120_000_000_000, PriceValid: 265,
			Description: "Тренажер, развивающий зрение до уровня 'всевидения'.",
		},
		{
			Name: "Импульс Квантовый прыжок", MaxLevel: 899, ValidCoef: 1.30, ClickCoef: 1.60, MinLeague: 8,
			PriceGrowthCoef: 1.25, PriceClick: 125_000_000_000, PriceValid: 265,
			Description: "Импульс, позволяющий совершить квантовый скачок в возможностях.",
		},
		{
			Name: "Кибернетический имплант Верховный разум", MaxLevel: 899, ValidCoef: 1.30, ClickCoef: 1.61, MinLeague: 8,
			PriceGrowthCoef: 1.25, PriceClick: 130_000_000_000, PriceValid: 270,
			Description: "Имплант, наделяющий сверхинтеллектом и интуицией.",
		},
		{
			Name: "Экзоскелет Агент-Феномен", MaxLevel: 899, ValidCoef: 1.30, ClickCoef: 1.62, MinLeague: 8,
			PriceGrowthCoef: 1.25, PriceClick: 135_000_000_000, PriceValid: 270,
			Description: "Экзоскелет, превращающий агента в феноменальную силу.",
		},
		{
			Name: "Сверхзвуковой протокол Гром", MaxLevel: 899, ValidCoef: 1.30, ClickCoef: 1.63, MinLeague: 8,
			PriceGrowthCoef: 1.25, PriceClick: 140_000_000_000, PriceValid: 275,
			Description: "Протокол, выводящий скорость реакции на сверхзвуковой уровень.",
		},
		{
			Name: "Финальный драйв Апогей", MaxLevel: 5, ValidCoef: 1.30, ClickCoef: 1.64, MinLeague: 8,
			PriceGrowthCoef: 1.25, PriceClick: 145_000_000_000, PriceValid: 275,
			Description: "Абсолютное улучшение, выводящее агента на пик возможностей.",
		},
	}

	for _, up := range updates {
		db.FirstOrCreate(&up, model.Update{Name: up.Name})
	}
}
