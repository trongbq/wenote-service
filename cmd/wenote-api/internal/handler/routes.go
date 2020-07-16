package handler

import (
	"net/http"
	"time"
	"wenote/cmd/wenote-api/internal/error"

	"github.com/gin-gonic/gin"
)

// ServiceHandler contains all handler of the app to be served under routes
type ServiceHandler struct {
	userHandler *UserHandler
	authHandler *AuthHandler
}

// NewServiceHandler creates new ServiceHandler
func NewServiceHandler(
	userHandler *UserHandler,
	authHandler *AuthHandler,
) *ServiceHandler {
	return &ServiceHandler{
		userHandler,
		authHandler,
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

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/v1")
	{
		v1User := v1.Group("users")
		{
			v1User.GET("/:id", handlers.userHandler.GetUserByID)
		}
		v1Auth := v1.Group("auth")
		{
			v1Auth.POST("/register", handlers.authHandler.Register)
			v1Auth.POST("/signin", handlers.authHandler.SignIn)
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
