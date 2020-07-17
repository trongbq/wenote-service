package main

import (
	"wenote/cmd/wenote-api/config"
	"wenote/cmd/wenote-api/internal/handler"
	"wenote/cmd/wenote-api/log"
	"wenote/internal/account"
	"wenote/internal/storage"
	"wenote/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {
	log.InitLogrus()
	config.LoadConfig()

	// Init storage
	storage, err := storage.NewStorage()
	if err != nil {
		panic(err)
	}

	// Init services
	userService := user.NewService(storage)
	accountService := account.NewService(userService, storage)

	// Init handlers
	userHandler := handler.NewUserHandler(userService)
	accountHandler := handler.NewAuthHandler(accountService)
	serviceHandler := handler.NewServiceHandler(userHandler, accountHandler)

	// Setup routes and server
	router := gin.Default()
	handler.Routes(router, serviceHandler)
	router.Run()
}
