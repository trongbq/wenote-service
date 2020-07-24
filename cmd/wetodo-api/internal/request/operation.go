package request

import (
	"wetodo/internal/operation"
)

type SaveOperationRequest struct {
	Operations []operation.Operation
}
