package main

import (
	"time"

	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Accept-Encoding", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")
	auth := api.Group("/auth")
	auth.POST("/signup", auth_handler.Signup)
	auth.POST("/signin", auth_handler.Signin)
	router.Run()
}
