package transaction

import (
	"reflect"
	"strings"
	"time"
	"wetodo/internal/storage"

	"github.com/sirupsen/logrus"
)

// Repository ...
type Repository interface {
	CreateOrUpdateTask(t storage.Task) (storage.Task, error)
	GetTaskByID(userID int, ID string) (storage.Task, bool)
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
func (s *Service) SaveTransactions(userID int, transactions []Transaction) []error {
	var errs []error

	for _, tr := range transactions {
		switch tr.Entity {
		case EntityTask:
			task, found := s.r.GetTaskByID(userID, tr.ID)
			if found == false {
				task.ID = tr.ID
				task.UserID = userID
			}
			for _, a := range tr.Actions {
				if a.Method == MethodSet {
					setTaskFieldValue(&task, a.Argument)
				}
			}

			if _, err := s.r.CreateOrUpdateTask(task); err != nil {
				errs = append(errs, SaveOperationError{task.ID, err.Error()})
			}
		default:
			logrus.Infof("Invalid entity: %v", tr.Entity)
		}
	}

	// Group operations by operation ID (task ID)
	// groups := groupByID(ops)

	// // Iterate through all groups
	// for _, ops := range groups {
	// 	var task storage.Task
	// 	var found bool
	// outer:
	// 	for _, op := range ops {
	// 		switch op.Type {
	// 		case AddTask:
	// 			task = op.Content.CopyToTask()
	// 			task.ID = op.ID
	// 			task.UserID = userID
	// 			task.CreatedAt = op.StartedAt
	// 		case UpdateTask:
	// 			if len(task.ID) == 0 {
	// 				if task, found = s.r.GetTaskByID(userID, op.ID); found == false {
	// 					errs = append(errs, TaskNotFoundError{op.ID})
	// 					break outer
	// 				}
	// 			}
	// 			// Update operation data to current task
	// 			task = op.Content.UpdateToTask(task)
	// 			task.UpdatedAt = op.StartedAt
	// 		case RemoveTask:
	// 			if len(task.ID) == 0 {
	// 				if task, found = s.r.GetTaskByID(userID, op.ID); found == false {
	// 					errs = append(errs, TaskNotFoundError{op.ID})
	// 					break outer
	// 				}
	// 			}
	// 			task.DeletedAt = ptrTime(time.Now())
	// 		case CompleteTask:
	// 			if len(task.ID) == 0 {
	// 				if task, found = s.r.GetTaskByID(userID, op.ID); found == false {
	// 					errs = append(errs, TaskNotFoundError{op.ID})
	// 					break outer
	// 				}
	// 			}
	// 			task.CompletedAt = ptrTime(time.Now())
	// 		default:
	// 			errs = append(errs, TypeError{op.ID, op.Type})
	// 		}
	// 	}

	// 	if len(task.ID) != 0 {
	// 		if _, err := s.r.CreateOrUpdateTask(task); err != nil {
	// 			errs = append(errs, SaveOperationError{task.ID, err.Error()})
	// 		}
	// 	}
	// }

	logrus.Debugf("Error: %v", errs)
	return errs
}

func ptrTime(t time.Time) *time.Time {
	return &t
}

func setTaskFieldValue(task *storage.Task, a Argument) {
	f := reflect.ValueOf(task).Elem().FieldByName(strings.Title(a.Name))
	t := f.Type().Name()
	if t == "string" {
		f.SetString(a.Value)
	}
}
