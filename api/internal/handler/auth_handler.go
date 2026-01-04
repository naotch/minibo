package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthService interface {
	Signup(email string, password string) error
}

type AuthHandler struct {
	service IAuthService
}

func NewAuthHandler(service IAuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Signup(c *gin.Context) {

	type input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	var request input
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Signup(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
