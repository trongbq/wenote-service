package transaction

import "fmt"

type UnmarshalError struct {
	err string
}

func (e UnmarshalError) Error() string {
	return fmt.Sprintf("Can not unmarshal request content: %v", e.err)
}

type TaskNotFoundError struct {
	ID string
}

func (e TaskNotFoundError) Error() string {
	return fmt.Sprintf("Task with ID: %v not found", e.ID)
}

type TypeError struct {
	ID   string
	Type string
}

func (e TypeError) Error() string {
	return fmt.Sprintf("Operation with ID: %v contains invalid type: %v", e.ID, e.Type)
}

type SaveOperationError struct {
	ID  string
	err string
}

func (e SaveOperationError) Error() string {
	return fmt.Sprintf("Operation with ID: %v can not be saved: %v", e.ID, e.err)
}
