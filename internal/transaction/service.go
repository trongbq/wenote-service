package transaction

import (
	"encoding/json"
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

	// Group transactions by ID
	groups := groupByEntityAndRecordID(transactions)

	// Iterate through all groups
	for entity, trs := range groups {
		switch entity {
		case entityTask:
			errs = append(errs, s.handleTaskTransactions(userID, trs)...)
		case entityChecklist:
		case entityTaskCategory:
		case entityTaskGoal:
		case entityTaskGroup:
		default:
			errs = append(errs, EntityTypeError{entity})
		}
	}

	logrus.Debugf("Error: %v", errs)
	return errs
}

// groupByEntityAndRecordID returns map of entity with values is map of record ID -> true transaction data
func groupByEntityAndRecordID(transactions []Transaction) map[string]map[string][]Transaction {
	groups := make(map[string]map[string][]Transaction)
	for _, t := range transactions {
		if _, ok := groups[t.Entity]; !ok {
			groups[t.Entity] = make(map[string][]Transaction)
		}
		groups[t.Entity][t.ID] = append(groups[t.Entity][t.ID], t)
	}
	return groups
}

func (s *Service) handleTaskTransactions(userID int, all map[string][]Transaction) []error {
	var errs []error

	for _, trs := range all {
		var task storage.Task
		var found bool
	outer:
		for _, tr := range trs {
			switch tr.Operation {
			case operationAdd:
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
			case operationUpdate:
				if len(task.ID) == 0 {
					if task, found = s.r.GetTaskByID(userID, tr.ID); found == false {
						errs = append(errs, RecordNotFoundError{entityTask, tr.ID})
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
			case operationDelete:
				if len(task.ID) == 0 {
					if task, found = s.r.GetTaskByID(userID, tr.ID); found == false {
						errs = append(errs, RecordNotFoundError{entityTask, tr.ID})
						break outer
					}
				}
				task.DeletedAt = ptrTime(time.Unix(int64(tr.At), 0))
			case operationComplete:
				if len(task.ID) == 0 {
					if task, found = s.r.GetTaskByID(userID, tr.ID); found == false {
						errs = append(errs, RecordNotFoundError{entityTask, tr.ID})
						break outer
					}
				}
				task.CompletedAt = ptrTime(time.Unix(int64(tr.At), 0))
			default:
				errs = append(errs, OperationError{tr.ID, tr.Operation})
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
