package main

import (
	"chat-back/database"
	"chat-back/server/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.GetDBInstance("main.db")
	r := gin.Default()

	r.Static("/static", "./static")
	r.StaticFile("/", "./static/index.html") // Отдача index.html по корневому пути "/"

	handler.RegisterControllers(r, db)

	r.Run(":8080")
}
