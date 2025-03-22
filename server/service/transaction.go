package service

import (
	"chat-back/database/model"
	"chat-back/database/repository"
	"fmt"

	"gorm.io/gorm"
)

type TransactionService struct {
	repo     repository.TransactionRepository
	userRepo repository.UserRepository
}

func NewTransactionService(db *gorm.DB) *TransactionService {
	repo := repository.NewTransactionRepository(db)
	userRepo := repository.NewUserRepository(db)

	return &TransactionService{repo: repo, userRepo: userRepo}
}

func (ts TransactionService) GetTransactionByUser(user *model.User) ([]model.Transaction, error) {
	transactions, err := ts.repo.FindByUser(user)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (ts TransactionService) GetById(transactionID uint) (*model.Transaction, error) {
	transaction, err := ts.repo.FindById(transactionID)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (ts *TransactionService) CreateTransaction(sender, receiver *model.User, sendValid, sendClick float64) (*model.Transaction, error) {
	if sender.UsualClicks < sendClick {
		return nil, fmt.Errorf("sender balance clicks less %v", sendClick)
	}

	if sender.ValidClicks < sendValid {
		return nil, fmt.Errorf("sender balance valid clicks less %v", sendValid)
	}

	newTransaction := model.Transaction{
		Sender:   *sender,
		Receiver: *receiver,
		Valid:    sendValid,
		Clicks:   sendClick,
	}
	err := ts.repo.Add(&newTransaction)
	if err != nil {
		return nil, err
	}

	sender.ValidClicks -= sendValid
	sender.UsualClicks -= sendClick

	receiver.ValidClicks += sendValid
	receiver.UsualClicks += sendClick

	errSender := ts.userRepo.Update(sender)
	errReceiver := ts.userRepo.Update(receiver)
	if errSender != nil || errReceiver != nil {
		return nil, err
	}

	return &newTransaction, nil
}
