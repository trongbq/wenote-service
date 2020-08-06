package main

import (
	"wetodo/cmd/wetodo-api/config"
	"wetodo/cmd/wetodo-api/internal/handler"
	"wetodo/cmd/wetodo-api/log"
	"wetodo/internal/account"
	"wetodo/internal/storage"
	"wetodo/internal/transaction"
	"wetodo/internal/user"

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
	accountService := account.NewService(storage)
	transactionService := transaction.NewService(storage)

	// Init handlers
	userHandler := handler.NewUserHandler(userService)
	accountHandler := handler.NewAuthHandler(accountService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	serviceHandler := handler.NewServiceHandler(userHandler, accountHandler, transactionHandler)

	// Setup routes and server
	router := gin.Default()
	handler.Routes(router, serviceHandler)
	err = router.Run()
	if err != nil {
		panic(err)
	}
}
