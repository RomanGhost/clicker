package repository

import (
	"chat-back/database/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Repository[model.User]
	FindByLogin(login string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindById(ID uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, ID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(object *model.User) error {
	return r.db.Save(object).Error
}

func (r *userRepository) Add(object *model.User) error {
	return r.db.Create(object).Error
}

func (r *userRepository) FindByLogin(login string) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, "login = ?", login).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
