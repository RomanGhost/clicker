package handler

import (
	"chat-back/database/model"
	"chat-back/server/jwtservice"
	"chat-back/server/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UpdateHandler struct {
	updateService *service.UpdateService
	userService   *service.UserService
}

func NewUpdateHandler(db *gorm.DB) *UpdateHandler {
	updateService := service.NewUpdateService(db)
	userService := service.NewUserService(db)
	return &UpdateHandler{
		updateService: updateService,
		userService:   userService,
	}
}

func (h *UpdateHandler) PostBuyUpdate(c *gin.Context) {
	var body struct {
		UpdateID uint `json:"update_id"`
	}
	// get cookies for auth
	tokenCookie, err := c.Request.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token in cookies not found"})
		return
	}
	tokenString := tokenCookie.Value
	parsedToken, err := jwtservice.GetFromJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to parse jwt"})
		c.Abort()
		return
	}

	user, err := h.userService.GetUserById(parsedToken.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User with id: %v not found", parsedToken.UserID)})
		log.Printf("User with id: %v not found", parsedToken.UserID)
		return
	}

	// read json
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request"})
		c.Abort()
		return
	}

	update, err := h.updateService.GetById(body.UpdateID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Update with id: %v not found", body.UpdateID)})
		log.Printf("Update with id: %v not found", body.UpdateID)
		return
	}

	var userUpdate *model.UserUpdate
	userUpdate, err = h.updateService.AddUpdateForUser(update, user)
	// если улучшение найдено
	if userUpdate != nil && err != nil {
		userUpdate, err = h.updateService.LevelupUpdateForUser(update, user)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Update error:%v", err)})
			log.Printf("Update error:%v", err)
			return
		}
	} else if err != nil {
		log.Printf("userUpdate error:%v", err)
		c.JSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("userUpdate error:%v", err)})
		return
	}

	// send info about update
	responceBody := struct {
		UpdateID  uint    `json:"update_id"`
		Level     uint    `json:"level"`
		ClickCoef float64 `json:"click_coef"`
		ValidCoef float64 `json:"valid_coef"`
	}{
		UpdateID:  userUpdate.UpdateID,
		Level:     userUpdate.Level,
		ClickCoef: userUpdate.ClickCoef,
		ValidCoef: userUpdate.ValidCoef,
	}
	c.JSON(http.StatusOK, responceBody)
}
