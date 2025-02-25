package handler

import (
	"chat-back/server/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TransactionHandler struct {
	service     *service.TransactionService
	userService *service.UserService
}

func NewTransactionHandler(db *gorm.DB) *TransactionHandler {
	transactionService := service.NewTransactionService(db)
	userService := service.NewUserService(db)
	return &TransactionHandler{service: transactionService, userService: userService}
}

func (h *TransactionHandler) PostCreateTransaction(c *gin.Context) {
	var body struct {
		SenderID            uint    `json:"sender_id"`
		ReceiverID          uint    `json:"receiver_id"`
		ClicksTransfer      float64 `json:"clicks"`
		ValidClicksTransfer float64 `json:"valid_clicks"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request"})
		c.Abort()
		return
	}

	senderUser, err := h.userService.GetUserById(body.SenderID)
	if err != nil {
		log.Printf("User sender not found, ID: %v, err: %v", body.SenderID, err)
		c.JSON(http.StatusNotFound, gin.H{"message": "Failed to get sender user"})
		c.Abort()
		return
	}

	receiverUser, err := h.userService.GetUserById(body.ReceiverID)
	if err != nil {
		log.Printf("User receiver not found, ID: %v, err: %v\n", body.SenderID, err)
		c.JSON(http.StatusNotFound, gin.H{"message": "Failed to get receiver user"})
		c.Abort()
		return
	}

	newTransaction, err := h.service.CreateTransaction(senderUser, receiverUser, body.ValidClicksTransfer, body.ClicksTransfer)
	if err != nil {
		log.Printf("Transaction error err: %v\n", err)
		c.JSON(http.StatusForbidden, gin.H{"message": fmt.Sprintf("error with transaction: %v", err)})
		c.Abort()
		return
	}

	sendBody := struct {
		ID            uint    `json:"id"`
		Valid         float64 `json:"valid_clicks"`
		Clicks        float64 `json:"clicks"`
		SenderLogin   string  `json:"sender_login"`
		ReceiverLogin string  `json:"receiver_login"`
	}{
		ID:            newTransaction.ID,
		Valid:         newTransaction.Valid,
		Clicks:        newTransaction.Clicks,
		SenderLogin:   newTransaction.Sender.Login,
		ReceiverLogin: newTransaction.Receiver.Login,
	}
	c.JSON(http.StatusOK, sendBody)
}
