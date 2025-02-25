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

func (s *UpdateService) LevelUpUpdateForUser(update *model.Update, user *model.User) error {
	userUpdate, err := s.userUpdateRepo.FindByUserUpdate(user, update)
	if err != nil {
		return fmt.Errorf("user update not found, err:%v", err)
	}
	if update.MaxLevel <= userUpdate.Level {
		return fmt.Errorf("max level")
	}

	priceClick := update.PriceClick
	priceValid := update.PriceValid
	for i := uint(0); i < userUpdate.Level; i++ {
		priceClick += priceClick * float64(update.PriceGrowthCoef)
		priceValid += priceValid * float64(update.PriceGrowthCoef)
	}

	if user.AllClicks < priceClick {
		return fmt.Errorf("user haven't clicks")
	}
	if user.ValidClicks < priceValid {
		return fmt.Errorf("user haven't valid clicks")
	}

	user.AllClicks -= priceClick
	user.ValidClicks -= priceValid

	err = s.userRepo.Update(user)
	if err != nil {
		return err
	}
	userUpdate.Level++
	err = s.userUpdateRepo.Update(userUpdate)
	if err != nil {
		return err
	}

	return nil
}
