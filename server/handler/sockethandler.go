package handler

import (
	"chat-back/database/model"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

type Message struct {
	TypeMessage string          `json:"typeMessage"`
	Data        json.RawMessage `json:"data"`
}

// Определяем структуру для пользователя
type Validate struct {
	Valid uint `json:"valid"`
	Nonce uint `json:"nonce"`
}

type UserSocketHandler struct {
	UserHandler
	upgrader  websocket.Upgrader
	clients   map[*websocket.Conn]*model.User
	broadcast chan struct{}
	mutex     *sync.Mutex
}

func NewUserSocketHandler(db *gorm.DB) *UserSocketHandler {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	var (
		clients   = make(map[*websocket.Conn]*model.User)
		broadcast = make(chan struct{})
		mutex     = &sync.Mutex{}
	)

	userHandler := *NewUserHandler(db)
	return &UserSocketHandler{
		UserHandler: userHandler,
		upgrader:    upgrader,
		clients:     clients,
		broadcast:   broadcast,
		mutex:       mutex,
	}
}

func (ush *UserSocketHandler) validMessage(valid, nonce uint, player *model.User, conn *websocket.Conn) error {
	// message format: "login_valid_nonce"
	sum := sha256.Sum256([]byte(fmt.Sprintf("%v_%v_%v", player.Login, valid, nonce)))
	log.Printf("Res of sum: %x\n", sum)
	if sum[0] != 0 && sum[1] < 128 {
		return fmt.Errorf("sha256 sum is not valid")
	}

	err := ush.service.ValidateMessage(valid, nonce, player)
	if err != nil {
		log.Printf("Error validate message: %v\n", err)
		return fmt.Errorf("error %v", err)
	}

	ush.clients[conn] = player
	return nil
}

func (ush *UserSocketHandler) HandleWebSocket(c *gin.Context) {
	var player *model.User

	// get cookies for auth
	tokenCookie, err := c.Request.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token in cookies not found"})
		return
	}
	tokenString := tokenCookie.Value

	// TODO: заменить на os.Getenv("SECRET")
	secret := []byte("testPhrase")
	parsedToken, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Error parsing"})
		log.Println("Ошибка при парсинге токена:", err)
		return
	}

	// Проверяем, что токен валиден и приводим claims к jwt.MapClaims:
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		// Извлекаем информацию:
		subVal := claims["sub"].(float64)

		userID := uint(subVal)
		player, err = ush.service.GetUserById(userID)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("User with id: %v not found", subVal)})
			log.Printf("User with id: %v not found", subVal)
			return
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
		log.Println("Token is invalid")
	}

	// обновляем соединение
	conn, err := ush.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "connection error"})
		log.Println("Ошибка при соединении:", err)
		return
	}

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
		// case "updateClicks":
		// 	ush.mutex.Lock()
		// 	err := ush.service.UpdateAllClicks(100, player)
		// 	ush.mutex.Unlock()
		// 	if err != nil {
		// 		continue
		// 	}
		case "valid":
			var validateMessage Validate
			if err := json.Unmarshal(message.Data, &validateMessage); err != nil {
				log.Fatalf("Ошибка при разборе данных: %v", err)
				continue
			}

			ush.mutex.Lock()
			err := ush.validMessage(validateMessage.Valid, validateMessage.Nonce, player, conn)
			ush.mutex.Unlock()
			if err != nil {
				continue
			}
		default:
			log.Println("Получено неизвестное сообщение:", message)
			continue
		}

		// Отправка обновлений всем клиентам
		select {
		case ush.broadcast <- struct{}{}:
		default:
			// Если канал заблокирован, пропускаем отправку
		}
	}
}

func (ush *UserSocketHandler) HandleMessages() {
	for range ush.broadcast {
		ush.mutex.Lock()
		scores := ""
		for _, player := range ush.clients {
			scores += fmt.Sprintf("%s: %d валидных кликов\n", player.Login, player.ValidClicks)
		}
		for client := range ush.clients {
			client.WriteMessage(websocket.TextMessage, []byte(scores))
		}
		ush.mutex.Unlock()
	}
}
