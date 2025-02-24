package repository

import (
	"chat-back/database/model"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) FindById(ID uint) (*model.Transaction, error) {
	var Transaction model.Transaction
	err := r.db.First(&Transaction, ID).Error
	if err != nil {
		return nil, err
	}
	return &Transaction, nil
}

func (r *TransactionRepository) Update(object *model.Transaction) error {
	return r.db.Save(object).Error
}

func (r *TransactionRepository) Add(object *model.Transaction) error {
	return r.db.Create(object).Error
}
