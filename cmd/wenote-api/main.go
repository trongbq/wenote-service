package main

import (
	"wenote/internal/auth"
	"wenote/internal/http/rest"
	"wenote/internal/http/rest/handler"
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
	serviceHandler := rest.NewServiceHandler(userHandler, authHandler)

	// Setup routes and server
	router := gin.Default()
	rest.Routes(router, serviceHandler)
	router.Run()
}
