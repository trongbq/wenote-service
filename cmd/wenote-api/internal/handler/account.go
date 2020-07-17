package handler

import (
	"net/http"
	"time"
	"wenote/cmd/wenote-api/internal/error"
	"wenote/cmd/wenote-api/internal/request"
	"wenote/cmd/wenote-api/internal/response"
	"wenote/internal/account"

	"github.com/gin-gonic/gin"
)

// AuthHandler is a handler for user resource
type AuthHandler struct {
	a *account.Service
}

// NewAuthHandler ...
func NewAuthHandler(a *account.Service) *AuthHandler {
	return &AuthHandler{a}
}

// Register register user and return token for authentication
func (h *AuthHandler) Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resp := error.Error{
			Code:      error.ErrorCodeBadRequest,
			Message:   err.Error(),
			Timestamp: time.Now(),
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	m := req.CopyToModel()
	auth, err := h.a.Register(m)
	if err != nil {
		switch err {
		case account.ErrFailedGenerateToken:
			c.JSON(http.StatusInternalServerError, error.Error{
				Code:      error.ErrorCodeInternalError,
				Message:   err.Error(),
				Timestamp: time.Now(),
			})
		default:
			c.JSON(http.StatusBadRequest, error.Error{
				Code:      error.ErrorCodeBadRequest,
				Message:   err.Error(),
				Timestamp: time.Now(),
			})
		}
		return
	}

	resp := response.CopyToAccountRegisterResponse(auth)
	c.JSON(http.StatusOK, resp)
}

// Login return token for authentication
func (h *AuthHandler) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resp := error.Error{
			Code:      error.ErrorCodeBadRequest,
			Message:   err.Error(),
			Timestamp: time.Now(),
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	auth, err := h.a.Login(req.Email, req.Password)
	if err != nil {
		switch err {
		case account.ErrUserNotFound:
			c.JSON(http.StatusNotFound, error.Error{
				Code:      error.ErrorCodeNotFound,
				Message:   err.Error(),
				Timestamp: time.Now(),
			})
		case account.ErrInvalidPassword:
			c.JSON(http.StatusBadRequest, error.Error{
				Code:      error.ErrorCodeBadRequest,
				Message:   err.Error(),
				Timestamp: time.Now(),
			})
		default:
			c.JSON(http.StatusInternalServerError, error.Error{
				Code:      error.ErrorCodeInternalError,
				Message:   err.Error(),
				Timestamp: time.Now(),
			})
		}
		return
	}

	resp := response.CopyToAccountRegisterResponse(auth)
	c.JSON(http.StatusOK, resp)
}

// Refresh return access token for authentication
func (h *AuthHandler) Refresh(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

// Logout return token for authentication
func (h *AuthHandler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
