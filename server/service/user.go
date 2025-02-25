package service

import (
	"chat-back/database/model"
	"chat-back/database/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
	userRepository := repository.NewUserRepository(db)
	return &UserService{repo: userRepository}
}

func (s *UserService) FindByLogin(login string) (*model.User, error) {
	return s.repo.FindByLogin(login)
}

func (s *UserService) RegisterUser(login, password string) (*model.User, error) {
	_, err := s.FindByLogin(login)
	if err == nil {
		return nil, fmt.Errorf("user exist")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	newUser := model.User{
		Login:       login,
		Password:    string(hashPassword),
		ValidClicks: 0,
		AllClicks:   0,
	}
	err = s.AddUser(&newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil

}

func (s *UserService) ComparePassword(user *model.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func (s *UserService) AddUser(user *model.User) error {
	return s.repo.Add(user)
}

func (s *UserService) GetUserById(userID uint) (*model.User, error) {
	return s.repo.FindById(userID)
}

func (s *UserService) UpdateAllClicks(countClicks float64, user *model.User) error {
	user.AllClicks += countClicks

	err := s.repo.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) ValidateMessage(valid, nonce float64, user *model.User, validCoef float64) error {
	if user.ValidClicks != valid {
		return fmt.Errorf("valid clicks invalid")
	}

	user.ValidClicks += 1
	// user.AllClicks += nonce

	err := s.repo.Update(user)
	if err != nil {
		return err
	}

	return nil
}
