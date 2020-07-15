package rest

import (
	"net/http"
	"time"
	"wenote/internal/http/rest/error"
	"wenote/internal/http/rest/handler"

	"github.com/gin-gonic/gin"
)

// ServiceHandler contains all handler of the app to be served under routes
type ServiceHandler struct {
	userHandler *handler.UserHandler
	authHandler *handler.AuthHandler
}

// NewServiceHandler creates new ServiceHandler
func NewServiceHandler(
	userHandler *handler.UserHandler,
	authHandler *handler.AuthHandler,
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

	apiV1 := router.Group("/api/v1")
	{
		apiV1User := apiV1.Group("users")
		{
			apiV1User.GET("/:id", handlers.userHandler.GetUserByID)
		}
		apiV1Auth := apiV1.Group("auth")
		{
			apiV1Auth.POST("/signup", handlers.authHandler.SignUp)
			apiV1Auth.POST("/signin", handlers.authHandler.SignIn)
		}
	}
	apiAdmin := router.Group("/api/admin/")
	{
		apiAdminUser := apiAdmin.Group("users")
		{
			apiAdminUser.GET("", handlers.userHandler.GetAllUsers)
		}
	}
}
