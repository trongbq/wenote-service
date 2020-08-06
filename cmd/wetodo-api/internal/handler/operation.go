package handler

import (
	"net/http"
	"reflect"
	"time"
	"wetodo/cmd/wetodo-api/internal/error"
	"wetodo/cmd/wetodo-api/internal/request"
	"wetodo/internal/operation"

	"github.com/gin-gonic/gin"
)

// OperationHandler handles operation actions
type OperationHandler struct {
	s *operation.Service
}

// NewOperationHandler return new OperationHandler
func NewOperationHandler(a *operation.Service) *OperationHandler {
	return &OperationHandler{a}
}

// SaveOperations handles persisting operation into storage
func (h *OperationHandler) SaveOperations(c *gin.Context) {
	var req request.SaveOperationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, error.SimpleBadRequestResponse(err.Error()))
		return
	}

	errs := h.s.SaveOperations(c.GetInt("UserID"), req.Operations)
	if len(errs) != 0 {
		info := make(map[string]interface{})
		for _, err := range errs {
			switch err.(type) {
			case operation.TaskNotFoundError:
				info[err.(operation.TaskNotFoundError).ID] = "TaskNotFoundError"
			case operation.TypeError:
				info[err.(operation.TypeError).ID] = "TypeError"
			case operation.SaveOperationError:
				info[err.(operation.SaveOperationError).ID] = "SaveOperationError"
			default:
				info[time.Now().String()] = reflect.TypeOf(err).String()
			}
		}
		err := error.Error{
			Code:      error.ErrorCodeUnknown,
			Message:   "Some issues with operations save",
			Info:      info,
			Timestamp: time.Now(),
		}
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
