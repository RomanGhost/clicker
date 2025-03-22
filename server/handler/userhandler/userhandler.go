package userhandler

import (
	"chat-back/server/service"

	"gorm.io/gorm"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	userService := service.NewUserService(db)
	userHandler := UserHandler{
		service: userService,
	}

	return &userHandler
}
