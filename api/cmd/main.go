package main

import (
	"github.com/gin-gonic/gin"
	"github.com/naotch/minibo/api/internal/handler"
	"github.com/naotch/minibo/api/internal/repository"
	"github.com/naotch/minibo/api/internal/service"
)

func main() {
	auth_repository := repository.NewAuthRepository(repository.DB)
	auth_service := service.NewAuthService(auth_repository)
	auth_handler := handler.NewAuthHandler(auth_service)

	router := gin.Default()
	auth := router.Group("/auth")
	auth.POST("/signup", auth_handler.Signup)
	router.Run()
}
