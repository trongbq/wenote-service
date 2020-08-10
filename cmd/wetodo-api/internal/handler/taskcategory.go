package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wetodo/internal/taskcategory"
)

type TaskCategoryHandler struct {
	s *taskcategory.Service
}

func NewTaskCategoryHandler(s *taskcategory.Service) *TaskCategoryHandler {
	return &TaskCategoryHandler{s}
}

func (h *TaskCategoryHandler) GetAllTaskCategoriesByUser(c *gin.Context) {
	c.JSON(http.StatusOK, h.s.GetAllTaskCategoriesByUser(c.GetInt("UserID")))
}
