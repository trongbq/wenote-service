package handler

import (
	"net/http"
	"wetodo/cmd/wetodo-api/internal/error"
	"wetodo/cmd/wetodo-api/internal/request"
	"wetodo/internal/operation"

	"github.com/gin-gonic/gin"
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

	errs := h.s.SaveOperations(c.GetInt("UserID"), req.Operations)
	if len(errs) != 0 {
		// TODO: handle error fully
		c.JSON(http.StatusBadRequest, error.SimpleBadRequestResponse(errs[0].Error()))
		return
	}

	c.JSON(http.StatusOK, nil)
}
