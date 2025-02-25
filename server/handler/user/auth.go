package user

import (
	"chat-back/database/model"
	"chat-back/server/jwtservice"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func generateToken(c *gin.Context, user *model.User) {
	token := jwtservice.JWTToken{
		UserID:    user.ID,
		UserLogin: user.Login,
		TimeLimit: time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	tokenString, err := token.ToString()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create session"})
		c.Abort()
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", *tokenString, 3600*32*1, "", "", false, false)

	c.JSON(http.StatusOK, gin.H{"message": "logged in successfully", "token": tokenString})
}

func (h *UserHandler) PostSignUp(c *gin.Context) {
	var body struct {
		Login          string `json:"login"`
		Password       string `json:"password"`
		AcceptPassword string `json:"accept_password"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request"})
		c.Abort()
		return
	}

	if body.Password != body.AcceptPassword {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Passwords didn't not match"})
		c.Abort()
		return
	}

	user, err := h.service.RegisterUser(body.Login, body.Password)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err})
		c.Abort()
		return
	}

	generateToken(c, user)
}

func (h *UserHandler) PostLogin(c *gin.Context) {
	var body struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request"})
		c.Abort()
		return
	}

	user, err := h.service.FindByLogin(body.Login)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err})
		c.Abort()
		return
	}

	result := h.service.ComparePassword(user, body.Password)
	if !result {
		c.JSON(http.StatusConflict, gin.H{"message": "Incorrect Password"})
		c.Abort()
		return
	}

	generateToken(c, user)
}

func Logout(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil || len(tokenString) <= 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cookie is not found"})
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, -1, "", "", false, false)
	c.JSON(http.StatusOK, gin.H{"message": "logout done"})
}
