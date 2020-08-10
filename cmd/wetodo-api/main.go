package main

import (
	"wetodo/cmd/wetodo-api/config"
	"wetodo/cmd/wetodo-api/internal/handler"
	"wetodo/cmd/wetodo-api/log"
	"wetodo/internal/account"
	"wetodo/internal/storage"
	"wetodo/internal/tag"
	"wetodo/internal/task"
	"wetodo/internal/taskcategory"
	"wetodo/internal/taskgoal"
	"wetodo/internal/transaction"
	"wetodo/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {
	log.InitLogrus()
	config.LoadConfig()

	// Init storage
	s, err := storage.NewStorage()
	if err != nil {
		panic(err)
	}

	// Init services
	userService := user.NewService(s)
	accountService := account.NewService(s)
	transactionService := transaction.NewService(s)
	categoryService := taskcategory.NewService(s)
	tagService := tag.NewService(s)
	taskGoalService := taskgoal.NewService(s)
	taskService := task.NewService(s)

	// Init handlers
	userHandler := handler.NewUserHandler(userService)
	accountHandler := handler.NewAuthHandler(accountService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	categoryHandler := handler.NewTaskCategoryHandler(categoryService)
	tagHandler := handler.NewTagHandler(tagService)
	taskGoalHandler := handler.NewTagGoalHandler(taskGoalService)
	taskHandler := handler.NewTaskHandler(taskService)

	serviceHandler := handler.NewServiceHandler(
		userHandler,
		accountHandler,
		transactionHandler,
		categoryHandler,
		taskGoalHandler,
		taskHandler,
		tagHandler)

	// Setup routes and server
	router := gin.Default()
	handler.Routes(router, serviceHandler)
	err = router.Run()
	if err != nil {
		panic(err)
	}
}
