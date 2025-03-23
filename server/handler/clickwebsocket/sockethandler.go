package clickwebsocket

import (
	"chat-back/database/model"
	"chat-back/server/jwtservice"
	"chat-back/server/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

type ClickSocketHandler struct {
	userService *service.UserService
	upgrader    websocket.Upgrader
	clients     map[*websocket.Conn]*model.User
	broadcast   chan Message
	mutex       *sync.Mutex
}

func NewClickSocketHandler(db *gorm.DB) *ClickSocketHandler {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	var (
		clients   = make(map[*websocket.Conn]*model.User)
		broadcast = make(chan Message)
		mutex     = &sync.Mutex{}
	)

	userService := service.NewUserService(db)
	return &ClickSocketHandler{
		userService: userService,
		upgrader:    upgrader,
		clients:     clients,
		broadcast:   broadcast,
		mutex:       mutex,
	}
}

func (ush *ClickSocketHandler) HandleWebSocket(c *gin.Context) {
	var player *model.User

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

	player, err = ush.userService.GetUserById(parsedToken.UserID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("User with id: %v not found", parsedToken.UserID)})
		log.Printf("User with id: %v not found", parsedToken.UserID)
		return
	}
	// update connection
	conn, err := ush.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "connection error"})
		log.Println("Ошибка при соединении:", err)
		return
	}

	userInfo := UserInfo{
		player.ValidClicks,
		player.UsualClicks,
	}

	userJson, err := json.Marshal(userInfo)
	if err != nil {
		log.Println("Error marshal user info")
	}
	userMessage := Message{TypeMessage: "user", Data: userJson}
	conn.WriteJSON(userMessage)

	closeFunction := func() {
		conn.Close()
		// Удаляем клиента из списка при завершении соединения
		ush.mutex.Lock()
		delete(ush.clients, conn)
		ush.mutex.Unlock()
	}
	defer closeFunction()

	// Добавляем пользователя в список клиентов
	ush.mutex.Lock()
	ush.clients[conn] = player
	ush.mutex.Unlock()

	// Обработка сообщений от клиента
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Ошибка при чтении сообщения:", err)
			break
		}

		if len(msg) == 0 {
			continue // Пропускаем пустые сообщения
		}

		var message Message
		err = json.Unmarshal(msg, &message)
		if err != nil {
			log.Printf("message invalid: %v \n", err)
			continue
		}

		switch message.TypeMessage {
		case "click_batch":
			var batchMessage ClickBatch
			err = json.Unmarshal(message.Data, &batchMessage)
			if err != nil {
				log.Printf("Error with read batch message: %v, json: %v", err, string(message.Data))
				continue
			}
			log.Printf("Get batch: %+v", batchMessage)
			updateClicks := ValidateBatch(&batchMessage)

			ush.mutex.Lock()
			//Update user score
			err = ush.userService.UpdateAllClicks(updateClicks, player)
			if err != nil {
				ush.mutex.Unlock()
				log.Printf("Error validate batch message: %v\n", err)
				continue
			}
			ush.mutex.Unlock()

		case "valid":
			var validateMessage Validate
			if err := json.Unmarshal(message.Data, &validateMessage); err != nil {
				log.Printf("Ошибка при разборе данных: %v", err)
				continue
			}
			messageValidErr := ValidateMessageValid(validateMessage, player.Login)
			if messageValidErr == nil {
				ush.mutex.Lock()
				// update user
				err = ush.userService.ValidateMessage(validateMessage.Valid, validateMessage.Nonce, player)
				if err != nil {
					ush.mutex.Unlock()
					log.Printf("Error validate message: %v\n", err)
					continue
				}
			} else {
				continue
			}

		default:
			log.Println("Получено неизвестное сообщение:", message)
			continue
		}

		userInfo := struct {
			ValidClicks float64 `json:"valid_clicks"`
			Clicks      float64 `json:"all_clicks"`
		}{
			player.ValidClicks,
			player.UsualClicks,
		}

		userJson, err := json.Marshal(userInfo)
		if err != nil {
			log.Println("Error marshal user info")
			continue
		}
		userMessage := Message{TypeMessage: "user", Data: userJson}
		conn.WriteJSON(userMessage)

		select {
		case ush.broadcast <- Message{TypeMessage: "info", Data: []byte("")}:
		default:
			// Если канал заблокирован, пропускаем отправку
		}
	}
}

func (ush *ClickSocketHandler) HandleMessages() {
	for range ush.broadcast {
		ush.mutex.Lock()
		scores := ""
		for _, player := range ush.clients {
			scores += fmt.Sprintf("%s: %v валидных кликов\n", player.Login, player.ValidClicks)
		}
		// for client := range ush.clients {
		// 	client.WriteMessage(websocket.TextMessage, []byte(scores))
		// }
		ush.mutex.Unlock()
	}
}
