package handler

import (
	"net/http"
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
		c.JSON(http.StatusBadRequest, error.SimpleBadRequestResponse(err.Error()))
		return
	}

	m := req.CopyToModel()
	auth, err := h.a.Register(m)
	if err != nil {
		switch err {
		case account.ErrFailedGenerateToken:
			c.JSON(http.StatusInternalServerError, error.SimpleInternalErrorResponse(err.Error()))
		default:
			c.JSON(http.StatusBadRequest, error.SimpleBadRequestResponse(err.Error()))
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
		c.JSON(http.StatusBadRequest, error.SimpleBadRequestResponse(err.Error()))
		return
	}

	auth, err := h.a.Login(req.Email, req.Password)
	if err != nil {
		switch err {
		case account.ErrUserNotFound:
			c.JSON(http.StatusNotFound, error.SimpleErrorResponse(error.ErrorCodeBadRequest, err.Error()))
		case account.ErrInvalidPassword:
			c.JSON(http.StatusBadRequest, error.SimpleBadRequestResponse(err.Error()))
		default:
			c.JSON(http.StatusInternalServerError, error.SimpleInternalErrorResponse(err.Error()))
		}
		return
	}

	resp := response.CopyToAccountRegisterResponse(auth)
	c.JSON(http.StatusOK, resp)
}

// Refresh returns access token for authentication
func (h *AuthHandler) Refresh(c *gin.Context) {
	var req request.RefreshOauthTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, error.SimpleBadRequestResponse(err.Error()))
		return
	}

	auth, err := h.a.RefreshAccessToken(req.RefreshToken)
	if err != nil {
		switch err {
		case account.ErrInvalidRefreshToken:
			c.JSON(http.StatusBadRequest, error.SimpleBadRequestResponse(err.Error()))
		default:
			c.JSON(http.StatusInternalServerError, error.SimpleInternalErrorResponse(err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, auth)
}

// Logout returns token for authentication
func (h *AuthHandler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
