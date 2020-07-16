package handler

import (
	"net/http"
	"time"
	"wenote/cmd/wenote-api/internal/error"
	"wenote/internal/auth"

	"github.com/gin-gonic/gin"
)

// AuthHandler is a handler for user resource
type AuthHandler struct {
	a *auth.Service
}

// SignUpRequest contains request data for SignUp handler
type SignUpRequest struct {
	Name     string
	Email    string
	Password string
}

// SignInRequest contains request data for SignIn handler
type SignInRequest struct {
	Email    string
	Password string
}

// NewAuthHandler ...
func NewAuthHandler(a *auth.Service) *AuthHandler {
	return &AuthHandler{a}
}

// Register register user and return token for authentication
func (h *AuthHandler) Register(c *gin.Context) {
	var s SignUpRequest
	err := c.ShouldBindJSON(&s)
	if err != nil {
		resp := error.Error{
			Code:      error.ErrorCodeBadRequest,
			Message:   err.Error(),
			Timestamp: time.Now(),
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	u, err := h.a.SignUp(s.Name, s.Email, s.Password)
	if err != nil {
		resp := error.Error{
			Code:      error.ErrorCodeBadRequest,
			Message:   err.Error(),
			Timestamp: time.Now(),
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, u)
}

// SignIn return token for authentication
func (h *AuthHandler) SignIn(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
