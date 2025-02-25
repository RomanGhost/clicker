package repository

import (
	"chat-back/database/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserUpdateRepository interface {
	Repository[model.UserUpdate] //interface
	FindByUserUpdate(user *model.User, update *model.Update) (*model.UserUpdate, error)
}

type userUpdateRepository struct {
	RepositoryStruct[model.UserUpdate]
}

func NewUserUpdateRepository(db *gorm.DB) UserUpdateRepository {
	return &userUpdateRepository{
		RepositoryStruct: RepositoryStruct[model.UserUpdate]{db: db},
	}
}

func (r *userUpdateRepository) FindByUserUpdate(user *model.User, update *model.Update) (*model.UserUpdate, error) {
	var userUpdate model.UserUpdate

	err := r.db.First(&userUpdate, "user_id = ? AND update_id = ?", user.ID, update.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Обработка случая, когда запись не найдена
			return nil, fmt.Errorf("database entry not found")
		}
		// Обработка других возможных ошибок
		return nil, err
	}
	return &userUpdate, nil
}
