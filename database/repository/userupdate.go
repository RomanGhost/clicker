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
	FindbyUser(user *model.User) ([]model.UserUpdate, error)
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

func (r *userUpdateRepository) FindbyUser(user *model.User) ([]model.UserUpdate, error) {
	var userUpdates []model.UserUpdate
	err := r.db.Find(&userUpdates, "user_id = ?", user.ID)
	if err != nil {
		return nil, fmt.Errorf("error with read userUpdates from db")
	}
	return userUpdates, nil
}
