package error

import (
	"fmt"
	"time"
)

const (
	ErrorCodeUnauthorized  = "UNAUTHORIZED"
	ErrorCodeBadRequest    = "BAD_REQUEST"
	ErrorCodeNotFound      = "NOT_FOUND"
	ErrorCodeNotImplement  = "NOT_IMPLEMENT"
	ErrorCodeInternalError = "INTERNAL_ERROR"
)

// Error struct for rest response
type Error struct {
	Code      string                 `json:"code"`
	Message   string                 `json:"message"`
	Info      map[string]interface{} `json:"info"`
	Timestamp time.Time              `json:"timestamp:`
}

func (err *Error) Error() string {
	return fmt.Sprintf("%s: %s", err.Code, err.Message)
}
