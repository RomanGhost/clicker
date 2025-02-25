package service

import (
	"chat-back/database/model"
	"chat-back/database/repository"
	"fmt"

	"gorm.io/gorm"
)

type UpdateService struct {
	repo           repository.Repository[model.Update]
	userUpdateRepo repository.UserUpdateRepository
	userRepo       repository.UserRepository
}

func NewUpdateService(db *gorm.DB) *UpdateService {
	repo := repository.NewRepository[model.Update](db)
	userUpdateRepo := repository.NewUserUpdateRepository(db)
	userRepo := repository.NewUserRepository(db)

	return &UpdateService{repo: repo, userUpdateRepo: userUpdateRepo, userRepo: userRepo}
}

func (s *UpdateService) GetById(ID uint) (*model.Update, error) {
	return s.repo.FindById(ID)
}

func (s *UpdateService) GetValidClickCoef(user *model.User) float64 {
	userUpdate, err := s.userUpdateRepo.FindbyUser(user)
	if err != nil {
		return 1
	}
	var resultCoef float64
	for _, usUp := range userUpdate {
		resultCoef += usUp.ValidCoef
	}
	return resultCoef
}

func (s *UpdateService) GetClickCoef(user *model.User) float64 {
	userUpdate, err := s.userUpdateRepo.FindbyUser(user)
	if err != nil {
		return 1
	}
	var resultCoef float64
	for _, usUp := range userUpdate {
		resultCoef += usUp.ClickCoef
	}
	return resultCoef
}

func (s *UpdateService) AddUpdateForUser(update *model.Update, user *model.User) (*model.UserUpdate, error) {
	userUpdate, err := s.userUpdateRepo.FindByUserUpdate(user, update)
	if userUpdate != nil && err == nil {
		return userUpdate, fmt.Errorf("userUpdate found")
	}

	if update.MinLeague > user.League.Number+1 {
		return nil, fmt.Errorf("user league is lower than required, User:%v, Update:%v", user.League.Number, update.MinLeague)
	}

	if user.AllClicks < update.PriceClick {
		return nil, fmt.Errorf("user clicks balance less update: %v", update.PriceClick)
	}

	if user.ValidClicks < update.PriceValid {
		return nil, fmt.Errorf("user valid balance less update: %v", update.PriceValid)
	}

	user.AllClicks -= update.PriceClick
	user.ValidClicks -= update.PriceValid

	err = s.userRepo.Update(user)
	if err != nil {
		return nil, fmt.Errorf("user update error")
	}

	newUserUpdate := model.UserUpdate{
		User:   *user,
		Update: *update,
		Level:  1,
	}

	newUserUpdate.ClickCoef *= update.ClickCoef
	newUserUpdate.ValidCoef *= update.ValidCoef

	err = s.userUpdateRepo.Add(&newUserUpdate)
	if err != nil {
		return nil, err
	}
	return &newUserUpdate, nil
}

func (s *UpdateService) LevelupUpdateForUser(update *model.Update, user *model.User) (*model.UserUpdate, error) {
	userUpdate, err := s.userUpdateRepo.FindByUserUpdate(user, update)
	if err != nil {
		return nil, fmt.Errorf("userUpdate not found, err:%v", err)
	}
	if update.MaxLevel <= userUpdate.Level {
		return nil, fmt.Errorf("max level")
	}

	priceClick := update.PriceClick
	priceValid := update.PriceValid
	for i := uint(1); i < userUpdate.Level; i++ {
		priceClick *= float64(update.PriceGrowthCoef)
		priceValid *= float64(update.PriceGrowthCoef)
	}

	if user.AllClicks < priceClick {
		return nil, fmt.Errorf("user haven't clicks")
	}
	if user.ValidClicks < priceValid {
		return nil, fmt.Errorf("user haven't valid clicks")
	}

	user.AllClicks -= priceClick
	user.ValidClicks -= priceValid

	err = s.userRepo.Update(user)
	if err != nil {
		return nil, err
	}
	userUpdate.Level++
	userUpdate.ClickCoef *= update.ClickCoef
	userUpdate.ValidCoef *= update.ValidCoef

	err = s.userUpdateRepo.Update(userUpdate)
	if err != nil {
		return nil, err
	}

	return userUpdate, nil
}
