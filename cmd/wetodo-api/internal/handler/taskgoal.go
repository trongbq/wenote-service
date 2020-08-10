package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wetodo/internal/taskgoal"
)

type TagGoalHandler struct {
	s *taskgoal.Service
}

func NewTagGoalHandler(s *taskgoal.Service) *TagGoalHandler {
	return &TagGoalHandler{s}
}

func (h *TagGoalHandler) GetTaskGoalByID(c *gin.Context) {
	tg, ok := h.s.GetTaskGoalByID(c.GetInt("UserID"), c.Param("id"))
	if !ok {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, tg)
}
