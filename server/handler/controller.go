package handler

import (
	"chat-back/server/handler/clickwebsocket"
	"chat-back/server/handler/userhandler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterControllers(r *gin.Engine, db *gorm.DB) {
	controllers := []func(*gin.Engine, *gorm.DB){
		registerAuthController,
		registerSocketController,
		registerTransactionController,
	}

	for _, controller := range controllers {
		controller(r, db)
	}
}

func registerAuthController(r *gin.Engine, db *gorm.DB) {
	uh := userhandler.NewUserHandler(db)

	r.POST("/signup", uh.PostSignUp)
	r.POST("/login", uh.PostLogin)
	r.POST("/logout", userhandler.Logout)
}

func registerSocketController(r *gin.Engine, db *gorm.DB) {
	ush := clickwebsocket.NewClickSocketHandler(db)

	r.GET("/ws", ush.HandleWebSocket)
	go ush.HandleMessages()
}

func registerTransactionController(r *gin.Engine, db *gorm.DB) {
	t := NewTransactionHandler(db)

	r.POST("/new_transaction", t.PostCreateTransaction)
}
