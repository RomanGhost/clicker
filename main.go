package main

import (
	"chat-back/database"
)

func main() {
	database.GetDBInstance("main.db")
	// r := gin.Default()

	// r.Static("/static", "./static")
	// r.StaticFile("/", "./static/index.html") // Отдача index.html по корневому пути "/"

	// handler.RegisterControllers(r, db)

	// r.Run(":8080")
}
