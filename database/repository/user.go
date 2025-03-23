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
	RepositoryStruct[model.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		RepositoryStruct: RepositoryStruct[model.User]{db: db},
	}
}
func (r *userRepository) FindByLogin(login string) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, "login = ?", login).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
