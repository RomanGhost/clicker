package handler

import (
	"chat-back/server/jwtservice"
	"chat-back/server/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateHandler struct {
	updateService *service.UpdateService
	userService   *service.UserService
}

func (h *UpdateHandler) PatchLevelUp(c *gin.Context) {
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

	var body struct {
		UpdateID uint `json:"update_id"`
	}

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

	err = h.updateService.LevelUpUpdateForUser(update, user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Update error:%v", err)})
		log.Printf("Update error:%v", err)
		return
	}
}
