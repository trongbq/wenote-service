package rest

import (
	"fmt"
	"time"
)

const (
	ERRCODE_UNAUTHORIZED    = "UNAUTHORIZED"
	ERRCODE_INVALID_REQUEST = "INVALID_REQUEST"
	ERRCODE_NOT_FOUND       = "NOT_FOUND"
	ERRCODE_NOT_IMPLEMENT   = "NOT_IMPLEMENT"
	ERRCODE_INTERNAL_ERROR  = "INTERNAL_ERROR"
)

type Error struct {
	Code      string                 `json:"code"`
	Message   string                 `json:"message"`
	Info      map[string]interface{} `json:"info"`
	Timestamp time.Time              `json:"timestamp:`
}

func (err *Error) Error() string {
	return fmt.Sprintf("%s: %s", err.Code, err.Message)
}
