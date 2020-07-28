package operation

import (
	"errors"
	"time"
)

var (
	ErrOperationType = errors.New("Invalid operation type")
	ErrContentFormat = errors.New("Error content format")
	ErrTaskNotFound  = errors.New("Error task not found")
)

// Repository ...
type Repository interface {
	CreateOrUpdateTask(t Task) (Task, error)
	GetTaskByID(userID int, ID string) (Task, bool)
}

// Service ...
type Service struct {
	r Repository
}

// NewService ...
func NewService(r Repository) *Service {
	return &Service{r}
}

// SaveOperations iterates through list of operation to persist data changes
// TODO: handle error when a task operations can not finished
func (s *Service) SaveOperations(userID int, ops []Operation) []error {
	var errs []error

	// Group operations by operation ID (task ID)
	groups := groupByID(ops)

	// Iterate through all groups
	for _, ops := range groups {
		var task Task
		var found bool
		for _, op := range ops {
			switch op.Type {
			case AddTask:
				task = op.Content.CopyToTask()
				task.ID = op.ID
				task.UserID = userID
				task.CreatedAt = op.StartedAt
			case UpdateTask:
				if len(task.ID) == 0 {
					if task, found = s.r.GetTaskByID(userID, op.ID); found == false {
						errs = append(errs, ErrTaskNotFound)
						break
					}
				}
				// Update operation data to current task
				task = op.Content.UpdateToTask(task)
				task.UpdatedAt = op.StartedAt
			case RemoveTask:
				if len(task.ID) == 0 {
					if task, found = s.r.GetTaskByID(userID, op.ID); found == false {
						errs = append(errs, ErrTaskNotFound)
						break
					}
				}
				task.Deleted = true
				task.DeletedAt = ptrTime(time.Now())
			case CompleteTask:
				if len(task.ID) == 0 {
					if task, found = s.r.GetTaskByID(userID, op.ID); found == false {
						errs = append(errs, ErrTaskNotFound)
						break
					}
				}
				task.Completed = true
				task.CompletedAt = ptrTime(time.Now())
			default:
				errs = append(errs, ErrOperationType)
			}
		}

		if _, err := s.r.CreateOrUpdateTask(task); err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

func groupByID(ops []Operation) map[string][]Operation {
	groups := make(map[string][]Operation)
	for _, op := range ops {
		groups[op.ID] = append(groups[op.ID], op)
	}
	return groups
}

func ptrTime(t time.Time) *time.Time {
	return &t
}
