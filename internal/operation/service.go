package operation

import (
	"encoding/json"
	"errors"
)

var (
	ErrOperationType = errors.New("Invalid operation type")
	ErrContentFormat = errors.New("Error content format")
)

// Repository ...
type Repository interface {
	MarkCompleteTaskByID(id int) error
	CreateTask(task Task) (Task, error)
}

// Service ...
type Service struct {
	r Repository
}

// NewService ...
func NewService(r Repository) *Service {
	return &Service{r}
}

// SaveOperations ...
// TODO: generate id with UUID on client side
func (s *Service) SaveOperations(userID int, ops []Operation) error {
	for _, op := range ops {
		switch op.Type {
		case AddTask:
			o := AddTaskOperation{}
			if err := json.Unmarshal([]byte(op.Content), &o); err != nil {
				return ErrContentFormat
			}

			t := o.CopyToTask()
			t.UserID = userID
			_, err := s.r.CreateTask(t)
			if err != nil {
				return err
			}
		case UpdateTask:
		case RemoveTask:
		case CompleteTask:
			o := CompleteTaskOperation{}
			if err := json.Unmarshal([]byte(op.Content), &o); err != nil {
				return ErrContentFormat
			}
			s.r.MarkCompleteTaskByID(o.ID)
		default:
			return ErrOperationType
		}
	}
	return nil
}
