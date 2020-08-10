package main

import (
	"wetodo/cmd/wetodo-api/config"
	"wetodo/cmd/wetodo-api/internal/handler"
	"wetodo/cmd/wetodo-api/log"
	"wetodo/internal/account"
	"wetodo/internal/storage"
	"wetodo/internal/tag"
	"wetodo/internal/taskcategory"
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
	categoryService := taskcategory.NewService(storage)
	tagService := tag.NewService(storage)

	// Init handlers
	userHandler := handler.NewUserHandler(userService)
	accountHandler := handler.NewAuthHandler(accountService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	categoryHandler := handler.NewTaskCategoryHandler(categoryService)
	tagHandler := handler.NewTagHandler(tagService)

	serviceHandler := handler.NewServiceHandler(
		userHandler,
		accountHandler,
		transactionHandler,
		categoryHandler,
		tagHandler)

	// Setup routes and server
	router := gin.Default()
	handler.Routes(router, serviceHandler)
	err = router.Run()
	if err != nil {
		panic(err)
	}
}
