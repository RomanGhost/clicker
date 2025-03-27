package user

import (
	"chat-back/server/jwtservice"
	"chat-back/server/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	userService := service.NewUserService(db)
	userHandler := UserHandler{
		service: userService,
	}

	return &userHandler
}

type UserInfo struct {
	ID          uint    `json:"id"`
	Login       string  `json:"login"`
	UsualClicks float64 `json:"usual_clicks"`
	ValidClicks float64 `json:"valid_clicks"`
	LeagueID    uint    `json:"league_id"`
	League      string  `json:"league_code"`
}

func (h *UserHandler) GetUserByLogin(c *gin.Context) {
	userLogin, ok := c.GetQuery("userlogin")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Param 'userlogin' not found"})
	}

	user, err := h.service.GetUserByLogin(userLogin)
	if err != nil || user == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}

	userInfo := UserInfo{
		ID:          user.ID,
		UsualClicks: user.UsualClicks,
		ValidClicks: user.ValidClicks,
		LeagueID:    user.LeagueID,
		League:      user.League.Code,
	}

	c.JSON(http.StatusOK, userInfo)
}

func (h *UserHandler) GetUserProfile(c *gin.Context) {
	tokenCookie, err := c.Request.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token in cookies not found"})
		return
	}
	tokenString := tokenCookie.Value

	parsedToken, err := jwtservice.GetFromJWT(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Failed to parse jwt"})
		return
	}

	user, err := h.service.GetUserById(parsedToken.UserID)
	if err != nil || user == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}

	userInfo := UserInfo{
		ID:          user.ID,
		Login:       user.Login,
		UsualClicks: user.UsualClicks,
		ValidClicks: user.ValidClicks,
		LeagueID:    user.LeagueID,
		League:      user.League.Code,
	}

	c.JSON(http.StatusOK, userInfo)

}
