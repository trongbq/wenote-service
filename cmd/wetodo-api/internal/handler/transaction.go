package handler

import (
	"net/http"
	"reflect"
	"time"
	"wetodo/cmd/wetodo-api/internal/error"
	"wetodo/cmd/wetodo-api/internal/request"
	"wetodo/internal/transaction"

	"github.com/gin-gonic/gin"
)

// TransactionHandler handles transaction actions
type TransactionHandler struct {
	s *transaction.Service
}

// NewTransactionHandler return new TransactionHandler
func NewTransactionHandler(a *transaction.Service) *TransactionHandler {
	return &TransactionHandler{a}
}

// SaveTransactions handles persisting transactions into storage
func (h *TransactionHandler) SaveTransactions(c *gin.Context) {
	var req request.SaveTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, error.SimpleBadRequestResponse(err.Error()))
		return
	}

	errs := h.s.SaveTransactions(c.GetInt("UserID"), req.Transactions)
	if len(errs) != 0 {
		info := make(map[string]interface{})
		for _, err := range errs {
			switch err.(type) {
			case transaction.RecordNotFoundError:
				info[err.(transaction.RecordNotFoundError).ID] = "RecordNotFoundError"
			case transaction.EntityTypeError:
				info[err.(transaction.EntityTypeError).ID] = "EntityTypeError"
			case transaction.OperationError:
				info[err.(transaction.OperationError).ID] = "OperationError"
			case transaction.SaveOperationError:
				info[err.(transaction.SaveOperationError).ID] = "SaveOperationError"
			default:
				info[time.Now().String()] = reflect.TypeOf(err).String()
			}
		}
		err := error.Error{
			Code:      error.ErrorCodeUnknown,
			Message:   "Some issues with transactions save",
			Info:      info,
			Timestamp: time.Now(),
		}
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
