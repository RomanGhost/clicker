package handler

import (
	"chat-back/server/handler/clickwebsocket"
	"chat-back/server/handler/userhandler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterControllers(r *gin.Engine, db *gorm.DB) {
	apiLink := r.Group("api/v1/clicker")
	controllers := []func(*gin.RouterGroup, *gorm.DB){
		registerAuthController,
		registerSocketController,
		registerTransactionController,
	}

	for _, controller := range controllers {
		controller(apiLink, db)
	}
}

func registerAuthController(r *gin.RouterGroup, db *gorm.DB) {
	uh := userhandler.NewUserHandler(db)

	rg := r.Group("/auth")

	rg.POST("/signup", uh.PostSignUp)
	rg.POST("/login", uh.PostLogin)
	rg.POST("/logout", userhandler.Logout)
}

func registerSocketController(r *gin.RouterGroup, db *gorm.DB) {
	ush := clickwebsocket.NewClickSocketHandler(db)

	r.GET("/ws", ush.HandleWebSocket)
	go ush.HandleMessages()
}

func registerTransactionController(r *gin.RouterGroup, db *gorm.DB) {
	t := NewTransactionHandler(db)

	rg := r.Group("/transaction")

	rg.POST("/new", t.PostCreateTransaction)
	rg.GET("/get/all", t.GetTransactionByUser)
	rg.GET("/get", t.GetTransactionById)
}
