package repository

import (
	"gorm.io/gorm"
)

type RepositoryStruct[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) *RepositoryStruct[T] {
	return &RepositoryStruct[T]{db: db}
}

func (r *RepositoryStruct[T]) FindById(ID uint) (*T, error) {
	var object T
	err := r.db.Session(&gorm.Session{}).First(&object, ID).Error
	if err != nil {
		return nil, err
	}
	return &object, nil
}

func (r *RepositoryStruct[T]) Update(object *T) error {
	return r.db.Save(object).Error
}

func (r *RepositoryStruct[T]) Add(object *T) error {
	return r.db.Create(object).Error
}
