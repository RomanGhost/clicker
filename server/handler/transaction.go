package handler

import (
	"chat-back/server/jwtservice"
	"chat-back/server/service"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type transactionJSON struct {
	ID            uint    `json:"id"`
	Valid         float64 `json:"valid_clicks"`
	Clicks        float64 `json:"clicks"`
	SenderLogin   string  `json:"sender_login"`
	ReceiverLogin string  `json:"receiver_login"`
}

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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to read request"})
		return
	}

	// get users by id
	senderUser, err := h.userService.GetUserById(body.SenderID)
	if err != nil {
		log.Printf("User sender not found, ID: %v, err: %v", body.SenderID, err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Failed to get sender user"})
		return
	}

	receiverUser, err := h.userService.GetUserById(body.ReceiverID)
	if err != nil {
		log.Printf("User receiver not found, ID: %v, err: %v\n", body.SenderID, err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Failed to get receiver user"})
		return
	}

	// create new transaction
	newTransaction, err := h.service.CreateTransaction(senderUser, receiverUser, body.ValidClicksTransfer, body.ClicksTransfer)
	if err != nil {
		log.Printf("Transaction error err: %v\n", err)
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": fmt.Sprintf("error with transaction: %v", err)})
		return
	}

	//send result of transaction
	sendBody := transactionJSON{
		ID:            newTransaction.ID,
		Valid:         newTransaction.Valid,
		Clicks:        newTransaction.Clicks,
		SenderLogin:   newTransaction.Sender.Login,
		ReceiverLogin: newTransaction.Receiver.Login,
	}
	c.JSON(http.StatusOK, sendBody)
}

func (h *TransactionHandler) GetTransactionById(c *gin.Context) {
	res, ok := c.GetQuery("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Param id not found"})
		return
	}
	ID, err := strconv.ParseUint(res, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"error": "id isn't number"})
	}

	transaction, err := h.service.GetById(uint(ID))
	if err != nil || transaction == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Transaction with id: %v not found", ID)})
	}

	sendBody := transactionJSON{
		ID:            transaction.ID,
		Valid:         transaction.Valid,
		Clicks:        transaction.Clicks,
		SenderLogin:   transaction.Sender.Login,
		ReceiverLogin: transaction.Receiver.Login,
	}
	c.JSON(http.StatusOK, sendBody)
}

func (h *TransactionHandler) GetTransactionByUser(c *gin.Context) {
	// get cookies for auth
	log.Println("Cookies", c.Request.Cookies())
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

	user, err := h.userService.GetUserById(parsedToken.UserID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("User with id: %v not found", parsedToken.UserID)})
		log.Printf("User with id: %v not found", parsedToken.UserID)
		return
	}

	transactions, err := h.service.GetTransactionByUser(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"error": "Transactions not found"})
		log.Printf("User with id: %v not found", parsedToken.UserID)
		return
	}

	var sendBody struct {
		Transactions []transactionJSON `json:"transactions"`
	}

	for _, transaction := range transactions {
		sendBody.Transactions = append(sendBody.Transactions, transactionJSON{
			ID:            transaction.ID,
			Valid:         transaction.Valid,
			Clicks:        transaction.Clicks,
			SenderLogin:   transaction.Sender.Login,
			ReceiverLogin: transaction.Receiver.Login,
		})
	}

	c.JSON(http.StatusOK, sendBody)
}
