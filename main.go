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
}
