package handler

import (
	"chat-back/database/model"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

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

func (h *UserSocketHandler) ServeHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func (h *UserSocketHandler) connectUser(login string, conn *websocket.Conn) *model.User {
	player, err := h.service.FindByLogin(login)
	if err != nil {
		player = &model.User{Login: login, ValidClicks: 0}
		h.service.AddUser(player)
	}
	//TODO Переделать на нормальный rest
	conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("UserInfo:%v_%v_%v", player.Login, player.AllClicks, player.ValidClicks)))
	h.clients[conn] = player
	return player
}

func (h *UserSocketHandler) validMessage(message string, player *model.User, conn *websocket.Conn) {
	player, err := h.service.ValidateMessage(message, player.ID)
	if err != nil {
		log.Printf("Error validate message: %v\n", err)
		return
	}

	h.clients[conn] = player
}

func (h *UserSocketHandler) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Ошибка соединения:", err)
		return
	}
	defer conn.Close()

	var player *model.User

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			h.mutex.Lock()
			delete(h.clients, conn)
			h.mutex.Unlock()
			break
		}

		message := string(msg)

		h.mutex.Lock()
		if player == nil {
			player = h.connectUser(message, conn)
		} else if message[:5] == "valid" {
			// TODO Если ошибка то обработать и отослать клиенту(игроку)
			h.validMessage(message[6:], player, conn)
		}
		h.broadcast <- struct{}{}
		h.mutex.Unlock()
	}
}

func (h *UserSocketHandler) HandleMessages() {
	for range h.broadcast {
		h.mutex.Lock()
		scores := ""
		for _, player := range h.clients {
			scores += fmt.Sprintf("%s: %d валидных кликов\n", player.Login, player.ValidClicks)
		}
		for client := range h.clients {
			client.WriteMessage(websocket.TextMessage, []byte(scores))
		}
		h.mutex.Unlock()
	}
}
