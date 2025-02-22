package handler

import (
	"chat-back/database/repository"
	"chat-back/server/service"

	"gorm.io/gorm"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := UserHandler{
		service: userService,
	}

	return &userHandler
}
