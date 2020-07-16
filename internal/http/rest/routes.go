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

	v1 := router.Group("/v1")
	{
		v1User := v1.Group("users")
		{
			v1User.GET("/:id", handlers.userHandler.GetUserByID)
		}
		v1Auth := v1.Group("auth")
		{
			v1Auth.POST("/signup", handlers.authHandler.SignUp)
			v1Auth.POST("/signin", handlers.authHandler.SignIn)
		}
	}
	admin := router.Group("/admin/")
	{
		adminUser := admin.Group("users")
		{
			adminUser.GET("", handlers.userHandler.GetAllUsers)
		}
	}
}
