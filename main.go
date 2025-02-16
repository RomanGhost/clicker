package main

import (
	"chat-back/database"
	"chat-back/database/repository"
	"chat-back/server/handler"
	"chat-back/server/service"
	"fmt"
	"net/http"
)

func main() {
	db := database.GetDBInstance("main.db")

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("/", userHandler.ServeHome)
	http.HandleFunc("/ws", userHandler.HandleConnections)

	go userHandler.HandleMessages()

	fmt.Println("Сервер запущен на :8080")
	http.ListenAndServe(":8081", nil)
}
