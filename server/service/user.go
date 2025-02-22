package service

import (
	"chat-back/database/model"
	"chat-back/database/repository"
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"

	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
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

func (s *UserService) addAllClicks(user *model.User, delta uint) (*model.User, error) {
	user.AllClicks += delta

	err := s.repo.Update(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) addValidClicks(user *model.User, delta uint) (*model.User, error) {
	user.ValidClicks += delta

	err := s.repo.Update(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) ValidateMessage(message string, userID uint) (*model.User, error) {
	// message format: "login_valid_nonce"
	user, err := s.repo.FindById(userID)
	if err != nil {
		log.Printf("Get user error: %v\n", err)
		return nil, err
	}

	sum := sha256.Sum256([]byte(message))
	log.Printf("Res of sum: %x\n", sum)
	if sum[0] != 0 {
		return nil, fmt.Errorf("sha256 sum is not valid")
	}

	parts := strings.Split(message, "_")
	if len(parts) >= 3 {
		_, valid, nonce := parts[0], parts[1], parts[2]
		validU64, err := strconv.ParseUint(valid, 10, 64)
		if err != nil {
			return nil, err
		}
		validU := uint(validU64)
		if user.ValidClicks != validU {
			user.ValidClicks = validU
			err = s.repo.Update(user)
			if err != nil {
				return nil, err
			}
		}

		nonceU64, err := strconv.ParseUint(nonce, 10, 64)
		if err != nil {
			return nil, err
		}

		user, err = s.addAllClicks(user, uint(nonceU64))
		if err != nil {
			return nil, err
		}
		user, err = s.addValidClicks(user, 1)
		if err != nil {
			return nil, err
		}

		return user, nil
	} else {
		return nil, fmt.Errorf("not enough elements after splitting the string")
	}

}
