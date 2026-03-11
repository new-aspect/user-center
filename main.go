package main

import (
	"log"
	"user-center/internal/handler"
	"user-center/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	repository.Init()

	r := gin.Default()

	r.GET("/api/health", handler.GetHealth)
	r.POST("/api/users/register", handler.PostUserRegister)
	r.POST("/api/users/login", handler.PostUserLogin)

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run service: %v", err)
	}
}
