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
		dbInstance.AutoMigrate(&model.User{})
	})
	return dbInstance
}
