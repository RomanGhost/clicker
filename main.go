package main

import (
	"chat-back/database"
	"chat-back/server/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.GetDBInstance("main.db")
	r := gin.Default()

	handler.RegisterControllers(r, db)

	r.Run(":8080")
	// userRepository := repository.NewUserRepository(db)
	// userService := service.NewUserService(userRepository)
	// userHandler := handler.NewUserSocketHandler(userService)

	// http.HandleFunc("/", userHandler.ServeHome)
	// http.HandleFunc("/ws", userHandler.HandleConnections)

	// go userHandler.HandleMessages()

	// fmt.Println("Сервер запущен на :8080")
	// http.ListenAndServe(":8080", nil)
}
