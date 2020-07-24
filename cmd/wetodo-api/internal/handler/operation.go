package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wetodo/cmd/wetodo-api/internal/error"
	"wetodo/cmd/wetodo-api/internal/request"
	"wetodo/internal/operation"
)

type OperationHandler struct {
	s *operation.Service
}

func NewOperationHandler(a *operation.Service) *OperationHandler {
	return &OperationHandler{a}
}

func (h *OperationHandler) SaveOperations(c *gin.Context) {
	var req request.SaveOperationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, error.SimpleBadRequestResponse(err.Error()))
		return
	}

	h.s.SaveOperations(req.Operations)

	c.JSON(http.StatusOK, nil)
}
