package repository

import (
	"chat-back/database/model"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Repository[model.Transaction]
	FindByUser(user *model.User) ([]model.Transaction, error)
}

type transactionRepository struct {
	RepositoryStruct[model.Transaction]
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		RepositoryStruct: RepositoryStruct[model.Transaction]{db: db.Preload("Sender").Preload("Receiver")},
	}
}

func (r *transactionRepository) FindByUser(user *model.User) ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := r.db.Session(&gorm.Session{}).Find(&transactions, "sender_ID = ? OR receiver_ID = ?", user.ID, user.ID).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
