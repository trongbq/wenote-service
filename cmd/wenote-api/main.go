package main

import (
	"wenote/cmd/wenote-api/internal/handler"
	"wenote/internal/auth"
	"wenote/internal/storage"
	"wenote/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init storage
	storage, err := storage.NewStorage()
	if err != nil {
		panic(err)
	}

	// Init services
	userService := user.NewService(storage)
	authService := auth.NewService(userService)

	// Init handlers
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(authService)
	serviceHandler := handler.NewServiceHandler(userHandler, authHandler)

	// Setup routes and server
	router := gin.Default()
	handler.Routes(router, serviceHandler)
	router.Run()
}
