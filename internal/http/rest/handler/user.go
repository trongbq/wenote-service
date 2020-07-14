package handler

import (
	"net/http"
	"wenote/internal/user"

	"github.com/gin-gonic/gin"
)

// UserHandler is a handler for user resource
type UserHandler struct {
	s *user.Service
}

// ListUserResponse ...
type ListUserResponse struct {
	List []user.User `json:"list"`
}

// NewUserHandler ...
func NewUserHandler(s *user.Service) *UserHandler {
	return &UserHandler{s}
}

// GetAllUsers handler
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users := h.s.GetAllUsers()

	c.JSON(http.StatusOK, ListUserResponse{List: users})
}
