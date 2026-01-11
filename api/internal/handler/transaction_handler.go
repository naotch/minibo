package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ITransactionService interface {
	Record(userID uint, title string, category int, amount int) error
}

type TransactionHandler struct {
	service ITransactionService
}

func NewTransactionHandler(service ITransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) Record(c *gin.Context) {
	type input struct {
		Title    string `json:"title" binding:"required,max=20"`
		Category int    `json:"category" binding:"oneof=0 1"`
		Amount   int    `json:"amount" binding:"required,gt=0"`
	}

	var request input
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	val := c.MustGet("userId")
	var userID uint
	if fuid, ok := val.(float64); ok {
		userID = uint(fuid)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "System error"})
		return
	}

	err := h.service.Record(userID, request.Title, request.Category, request.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
}
