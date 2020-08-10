package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wetodo/internal/tag"
)

type TagHandler struct {
	s *tag.Service
}

func NewTagHandler(s *tag.Service) *TagHandler {
	return &TagHandler{s}
}

func (h *TagHandler) GetAllTagsByUser(c *gin.Context) {
	c.JSON(http.StatusOK, h.s.GetAllTagsByUser(c.GetInt("UserID")))
}
