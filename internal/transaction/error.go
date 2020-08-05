package transaction

import "fmt"

type UnmarshalError struct {
	err string
}

func (e UnmarshalError) Error() string {
	return fmt.Sprintf("Can not unmarshal request content: %v", e.err)
}

type RecordNotFoundError struct {
	entity string
	id     string
}

func (e RecordNotFoundError) Error() string {
	return fmt.Sprintf("%v with ID: %v not found", e.entity, e.id)
}

type EntityTypeError struct {
	entity string
}

func (e EntityTypeError) Error() string {
	return fmt.Sprintf("Invalid entity type: %v", e.entity)
}

type OperationError struct {
	id        string
	operation string
}

func (e OperationError) Error() string {
	return fmt.Sprintf("Invalid operation: %v on record: %v", e.operation, e.id)
}

type SaveOperationError struct {
	ID  string
	err string
}

func (e SaveOperationError) Error() string {
	return fmt.Sprintf("Operation with ID: %v can not be saved: %v", e.ID, e.err)
}
