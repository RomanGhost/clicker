package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterControllers(r *gin.Engine, db *gorm.DB) {
	registerAuthControler(r, db)
}

func registerAuthControler(r *gin.Engine, db *gorm.DB) {
	uh := NewUserHandler(db)

	r.POST("/signup", uh.PostSignUp)
	r.POST("/login", uh.PostLogin)
	r.POST("/logout", Logout)
}
