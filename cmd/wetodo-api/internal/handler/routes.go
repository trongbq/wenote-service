package handler

import (
	"net/http"
	"time"
	"wetodo/cmd/wetodo-api/internal/error"
	"wetodo/cmd/wetodo-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

// ServiceHandler contains all handler of the app to be served under routes
type ServiceHandler struct {
	userHandler         *UserHandler
	accountHandler      *AuthHandler
	transactionHandler  *TransactionHandler
	taskCategoryHandler *TaskCategoryHandler
	tagHandler          *TagHandler
}

// NewServiceHandler creates new ServiceHandler
func NewServiceHandler(
	userHandler *UserHandler,
	accountHandler *AuthHandler,
	transactionHandler *TransactionHandler,
	taskCategoryHandler *TaskCategoryHandler,
	tagHandler *TagHandler,
) *ServiceHandler {
	return &ServiceHandler{
		userHandler,
		accountHandler,
		transactionHandler,
		taskCategoryHandler,
		tagHandler,
	}
}

// Routes setup routes with handler
func Routes(router *gin.Engine, handlers *ServiceHandler) {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, error.Error{
			Code:      error.ErrorCodeNotFound,
			Message:   "Page not found",
			Timestamp: time.Now(),
		})
	})

	// Use to check whether client can connect to server API (in case client is offline)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/v1")
	{
		v1Auth := v1.Group("auth")
		{
			v1Auth.POST("/register", handlers.accountHandler.Register)
			v1Auth.POST("/login", handlers.accountHandler.Login)
			v1Auth.POST("/refresh", handlers.accountHandler.Refresh)
		}
		v1Auth.Use(middleware.AuthenticationMiddleware())
		{
			v1Auth.POST("/logout", handlers.accountHandler.Logout)
		}

		v1User := v1.Group("users")
		v1User.Use(middleware.AuthenticationMiddleware())
		{
			v1User.GET("/:id", handlers.userHandler.GetUserByID)
		}

		v1Transaction := v1.Group("transactions")
		v1Transaction.Use(middleware.AuthenticationMiddleware())
		{
			v1Transaction.POST("/save", handlers.transactionHandler.SaveTransactions)
		}

		v1Category := v1.Group("categories")
		v1Category.Use(middleware.AuthenticationMiddleware())
		{
			v1Category.GET("/", handlers.taskCategoryHandler.GetAllTaskCategoriesByUser)
		}

		v1Tags := v1.Group("tags")
		v1Tags.Use(middleware.AuthenticationMiddleware())
		{
			v1Tags.GET("/", handlers.tagHandler.GetAllTagsByUser)
		}
	}
	adminV1 := router.Group("/admin/v1")
	{
		adminV1User := adminV1.Group("users")
		{
			adminV1User.GET("", handlers.userHandler.GetAllUsers)
		}
	}
}
