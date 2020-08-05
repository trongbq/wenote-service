package transaction

import (
	"encoding/json"
	"time"
	"wetodo/internal/storage"

	"github.com/sirupsen/logrus"
)

// Repository ...
// TODO: to improve security, add check for userID (and taskID) for storage operation on table doesn't have userID column
type Repository interface {
	CreateOrUpdateTask(t storage.Task) (storage.Task, error)
	GetTaskByID(userID int, ID string) (storage.Task, bool)
	CreateOrUpdateChecklist(cl storage.Checklist) (storage.Checklist, error)
	GetChecklistByID(id string) (storage.Checklist, bool)
	DeleteChecklistByID(id string)
	CreateOrUpdateTaskCategory(tc storage.TaskCategory) (storage.TaskCategory, error)
	GetTaskCategoryByID(userID int, id string) (storage.TaskCategory, bool)
	CreateOrUpdateTaskGoal(tg storage.TaskGoal) (storage.TaskGoal, error)
	GetTaskGoalByID(userID int, id string) (storage.TaskGoal, bool)
	CreateOrUpdateTaskGroup(tg storage.TaskGroup) (storage.TaskGroup, error)
	GetTaskGroupByID(userID int, id string) (storage.TaskGroup, bool)
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
// Some transactions might depends on existing of other transactions so we have to save sequentially
func (s *Service) SaveTransactions(userID int, transactions []Transaction) []error {
	var errs []error

	// Iterate through all groups
	for _, tr := range transactions {
		switch tr.Entity {
		case entityTask:
			if err := s.handleTaskTransactions(userID, tr); err != nil {
				errs = append(errs, err)
			}
		case entityChecklist:
			if err := s.handleChecklistTransactions(tr); err != nil {
				errs = append(errs, err)
			}
		case entityTaskCategory:
		case entityTaskGoal:
		case entityTaskGroup:
		default:
			errs = append(errs, EntityTypeError{tr.ID, tr.Entity})
		}
	}

	logrus.Debugf("Error: %v", errs)
	return errs
}

func (s *Service) handleTaskTransactions(userID int, tr Transaction) error {
	var task storage.Task
	var found bool

	switch tr.Operation {
	case operationAdd:
		task.ID = tr.ID
		task.UserID = userID
		task.CreatedAt = time.Unix(int64(tr.At), 0)
		// Copy request content to task
		tc := TaskContent{}
		err := json.Unmarshal(tr.Args, &tc)
		if err != nil {
			return UnmarshalError{err.Error()}
		}
		tc.CopyToTask(&task)
	case operationUpdate:
		if len(task.ID) == 0 {
			if task, found = s.r.GetTaskByID(userID, tr.ID); found == false {
				return RecordNotFoundError{entityTask, tr.ID}
			}
		}
		// Update request data to current task
		tc := TaskContent{}
		err := json.Unmarshal(tr.Args, &tc)
		if err != nil {
			return UnmarshalError{err.Error()}
		}
		tc.CopyToTask(&task)
		task.UpdatedAt = time.Unix(int64(tr.At), 0)
	case operationDelete:
		if len(task.ID) == 0 {
			if task, found = s.r.GetTaskByID(userID, tr.ID); found == false {
				return RecordNotFoundError{entityTask, tr.ID}
			}
		}
		task.DeletedAt = ptrTime(time.Unix(int64(tr.At), 0))
	case operationComplete:
		if len(task.ID) == 0 {
			if task, found = s.r.GetTaskByID(userID, tr.ID); found == false {
				return RecordNotFoundError{entityTask, tr.ID}
			}
		}
		task.CompletedAt = ptrTime(time.Unix(int64(tr.At), 0))
	default:
		return OperationError{tr.ID, tr.Operation}
	}

	if len(task.ID) != 0 {
		if _, err := s.r.CreateOrUpdateTask(task); err != nil {
			return SaveOperationError{task.ID, err.Error()}
		}
	}

	return nil
}

func (s *Service) handleChecklistTransactions(tr Transaction) error {
	var checklist storage.Checklist
	var found bool

	switch tr.Operation {
	case operationAdd:
		checklist.ID = tr.ID
		checklist.CreatedAt = time.Unix(int64(tr.At), 0)
		// Copy request content to task
		clc := ChecklistContent{}
		err := json.Unmarshal(tr.Args, &clc)
		if err != nil {
			return UnmarshalError{err.Error()}
		}
		clc.CopyToChecklist(&checklist)
	case operationUpdate:
		if len(checklist.ID) == 0 {
			if checklist, found = s.r.GetChecklistByID(tr.ID); found == false {
				return RecordNotFoundError{entityTask, tr.ID}
			}
		}
		clc := ChecklistContent{}
		err := json.Unmarshal(tr.Args, &clc)
		if err != nil {
			return UnmarshalError{err.Error()}
		}
		clc.CopyToChecklist(&checklist)
		checklist.UpdatedAt = time.Unix(int64(tr.At), 0)
	case operationDelete:
		// Deletion is the end action of a checklist,
		// Remove record in DB or just breakout of loop for this checklist ID
		s.r.DeleteChecklistByID(tr.ID)
		return nil
	case operationComplete:
		if len(checklist.ID) == 0 {
			if checklist, found = s.r.GetChecklistByID(tr.ID); found == false {
				return RecordNotFoundError{entityTask, tr.ID}
			}
		}
		checklist.CompletedAt = ptrTime(time.Unix(int64(tr.At), 0))
	default:
		return OperationError{tr.ID, tr.Operation}
	}

	if len(checklist.ID) != 0 {
		if _, err := s.r.CreateOrUpdateChecklist(checklist); err != nil {
			return SaveOperationError{checklist.ID, err.Error()}
		}
	}

	return nil
}
