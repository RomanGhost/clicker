package handler

import (
	"chat-back/server/handler/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterControllers(r *gin.Engine, db *gorm.DB) {
	registerAuthControler(r, db)
	registerSocketController(r, db)
	registerTransactionController(r, db)
}

func registerAuthControler(r *gin.Engine, db *gorm.DB) {
	uh := user.NewUserHandler(db)

	r.POST("/signup", uh.PostSignUp)
	r.POST("/login", uh.PostLogin)
	r.POST("/logout", user.Logout)
}

func registerSocketController(r *gin.Engine, db *gorm.DB) {
	ush := user.NewUserSocketHandler(db)

	r.GET("/ws", ush.HandleWebSocket)
	go ush.HandleMessages()
}

func registerTransactionController(r *gin.Engine, db *gorm.DB) {
	t := NewTransactionHandler(db)

	r.POST("/new_transaction", t.PostCreateTransaction)
}
