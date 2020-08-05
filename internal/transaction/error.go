package transaction

import "fmt"

type UnmarshalError struct {
	err string
}

func (e UnmarshalError) Error() string {
	return fmt.Sprintf("Can not unmarshal request content: %v", e.err)
}

type RecordNotFoundError struct {
	Entity string
	ID     string
}

func (e RecordNotFoundError) Error() string {
	return fmt.Sprintf("%v with ID: %v not found", e.Entity, e.ID)
}

type EntityTypeError struct {
	ID     string
	Entity string
}

func (e EntityTypeError) Error() string {
	return fmt.Sprintf("Invalid entity type: %v on record: %v", e.Entity, e.ID)
}

type OperationError struct {
	ID        string
	Operation string
}

func (e OperationError) Error() string {
	return fmt.Sprintf("Invalid operation: %v on record: %v", e.Operation, e.ID)
}

type SaveOperationError struct {
	ID  string
	err string
}

func (e SaveOperationError) Error() string {
	return fmt.Sprintf("Operation with ID: %v can not be saved: %v", e.ID, e.err)
}
