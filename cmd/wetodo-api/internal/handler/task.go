package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wetodo/cmd/wetodo-api/internal/error"
	"wetodo/cmd/wetodo-api/internal/request"
	"wetodo/internal/task"
)

type TaskHandler struct {
	s *task.Service
}

func NewTaskHandler(s *task.Service) *TaskHandler {
	return &TaskHandler{s}
}

func (h *TaskHandler) GetAllTaskByUser(c *gin.Context) {
	var req request.TaskListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, error.SimpleBadRequestResponse(err.Error()))
		return
	}

	tasks := h.s.GetAllTasksByUser(c.GetInt("userID"), req.Completed, req.Page, req.Limit, req.SortField, req.SortOrder)
	c.JSON(http.StatusOK, tasks)
}
