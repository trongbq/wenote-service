package handler

import (
	"net/http"
	"wetodo/cmd/wetodo-api/internal/response"
	"wetodo/internal/tag"

	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	s *tag.Service
}

func NewTagHandler(s *tag.Service) *TagHandler {
	return &TagHandler{s}
}

func (h *TagHandler) GetAllTagsByUser(c *gin.Context) {
	tags := h.s.GetAllTagsByUser(c.GetInt("UserID"))
	c.JSON(http.StatusOK, response.CopyToTagListResponse(tags))
}
