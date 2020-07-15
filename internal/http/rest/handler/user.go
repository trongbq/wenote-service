package handler

import (
	"net/http"
	"strconv"
	"time"
	"wenote/internal/http/rest/error"
	"wenote/internal/user"

	"github.com/gin-gonic/gin"
)

// UserHandler is a handler for user resource
type UserHandler struct {
	s *user.Service
}

// ListUserResponse ...
type ListUserResponse struct {
	List []UserDetailResponse `json:"list"`
}

type UserDetailResponse struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	PictureURL string    `json:"picture_url"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

// NewUserHandler ...
func NewUserHandler(s *user.Service) *UserHandler {
	return &UserHandler{s}
}

// GetAllUsers handler
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	all := h.s.GetAllUsers()
	users := make([]UserDetailResponse, 0, len(all))
	for _, u := range all {
		userResp := UserDetailResponse{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		}
		users = append(users, userResp)
	}

	c.JSON(http.StatusOK, ListUserResponse{List: users})
}

// GetUserByID return user matches ID
func (h *UserHandler) GetUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp := error.Error{
			Code:      error.ErrorCodeBadRequest,
			Message:   "Can not parse user ID",
			Timestamp: time.Now(),
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	user, found := h.s.GetUserByID(userID)
	if found != true {
		resp := error.Error{
			Code:      error.ErrorCodeNotFound,
			Message:   "User not found",
			Timestamp: time.Now(),
		}
		c.JSON(http.StatusNotFound, resp)
		return
	}
	userResp := UserDetailResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	c.JSON(http.StatusOK, userResp)
}
