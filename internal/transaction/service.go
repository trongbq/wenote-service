package transaction

import (
	"encoding/json"
	"time"
	"wetodo/internal/storage"

	"github.com/sirupsen/logrus"
)

const (
	typeTaskAdd      = "TASK_ADD"
	typeTaskUpdate   = "TASK_UPDATE"
	typeTaskDelete   = "TASK_DELETE"
	typeTaskComplete = "TASK_COMPLETE"
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

	// Group transactions by ID
	groups := groupByID(transactions)

	// Iterate through all groups
	for _, trs := range groups {
		var task storage.Task
		var found bool
	outer:
		for _, tr := range trs {
			switch tr.Type {
			case typeTaskAdd:
				task.ID = tr.ID
				task.UserID = userID
				task.CreatedAt = time.Unix(int64(tr.At), 0)
				// Copy request content to task
				tc := TaskContent{}
				err := json.Unmarshal(tr.Args, &tc)
				if err != nil {
					errs = append(errs, UnmarshalError{err.Error()})
				}
				tc.CopyToTask(&task)
			case typeTaskUpdate:
				if len(task.ID) == 0 {
					if task, found = s.r.GetTaskByID(userID, tr.ID); found == false {
						errs = append(errs, TaskNotFoundError{tr.ID})
						break outer
					}
				}
				// Update request data to current task
				tc := TaskContent{}
				err := json.Unmarshal(tr.Args, &tc)
				if err != nil {
					errs = append(errs, UnmarshalError{err.Error()})
				}
				tc.CopyToTask(&task)
				task.UpdatedAt = time.Unix(int64(tr.At), 0)
			case typeTaskDelete:
				if len(task.ID) == 0 {
					if task, found = s.r.GetTaskByID(userID, tr.ID); found == false {
						errs = append(errs, TaskNotFoundError{tr.ID})
						break outer
					}
				}
				task.DeletedAt = ptrTime(time.Unix(int64(tr.At), 0))
			case typeTaskComplete:
				if len(task.ID) == 0 {
					if task, found = s.r.GetTaskByID(userID, tr.ID); found == false {
						errs = append(errs, TaskNotFoundError{tr.ID})
						break outer
					}
				}
				task.CompletedAt = ptrTime(time.Unix(int64(tr.At), 0))
			default:
				errs = append(errs, TypeError{tr.ID, tr.Type})
			}
		}

		if len(task.ID) != 0 {
			if _, err := s.r.CreateOrUpdateTask(task); err != nil {
				errs = append(errs, SaveOperationError{task.ID, err.Error()})
			}
		}
	}

	logrus.Debugf("Error: %v", errs)
	return errs
}

func groupByID(transactions []Transaction) map[string][]Transaction {
	groups := make(map[string][]Transaction)
	for _, t := range transactions {
		groups[t.ID] = append(groups[t.ID], t)
	}
	return groups
}
