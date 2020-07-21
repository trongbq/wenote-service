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

// SimpleErrorResponse builds simple format for error response
func SimpleErrorResponse(errorCode string, message string) Error {
	return Error{
		Code:      errorCode,
		Message:   message,
		Timestamp: time.Now(),
	}
}

// SimpleBadRequestResponse builds simple format for error response
func SimpleBadRequestResponse(message string) Error {
	return Error{
		Code:      ErrorCodeBadRequest,
		Message:   message,
		Timestamp: time.Now(),
	}
}

// SimpleInternalErrorResponse builds simple format for error response
func SimpleInternalErrorResponse(message string) Error {
	return Error{
		Code:      ErrorCodeInternalError,
		Message:   message,
		Timestamp: time.Now(),
	}
}
