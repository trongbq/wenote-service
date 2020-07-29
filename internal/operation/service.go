package operation

import (
	"time"
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
func (s *Service) SaveOperations(userID int, ops []Operation) []error {
	var errs []error

	// Group operations by operation ID (task ID)
	groups := groupByID(ops)

	// Iterate through all groups
	for _, ops := range groups {
		var task Task
		var found bool
	outer:
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
						errs = append(errs, TaskNotFoundError{op.ID})
						break outer
					}
				}
				// Update operation data to current task
				task = op.Content.UpdateToTask(task)
				task.UpdatedAt = op.StartedAt
			case RemoveTask:
				if len(task.ID) == 0 {
					if task, found = s.r.GetTaskByID(userID, op.ID); found == false {
						errs = append(errs, TaskNotFoundError{op.ID})
						break outer
					}
				}
				task.DeletedAt = ptrTime(time.Now())
			case CompleteTask:
				if len(task.ID) == 0 {
					if task, found = s.r.GetTaskByID(userID, op.ID); found == false {
						errs = append(errs, TaskNotFoundError{op.ID})
						break outer
					}
				}
				task.CompletedAt = ptrTime(time.Now())
			default:
				errs = append(errs, TypeError{op.ID, op.Type})
			}
		}

		if len(task.ID) != 0 {
			if _, err := s.r.CreateOrUpdateTask(task); err != nil {
				errs = append(errs, SaveOperationError{task.ID, err.Error()})
			}
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
