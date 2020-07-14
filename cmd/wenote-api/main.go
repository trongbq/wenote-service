package main

import (
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

	// Init handlers
	userHandler := handler.NewUserHandler(userService)
	serviceHandler := rest.NewServiceHandler(userHandler)

	// Setup routes and server
	router := gin.Default()
	rest.Routes(router, serviceHandler)
	router.Run()
}
