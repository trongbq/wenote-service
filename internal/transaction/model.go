package transaction

import (
	"encoding/json"
	"time"
	"wetodo/internal/storage"
)

const (
	entityTask         = "TASK"
	entityChecklist    = "CHECKLIST"
	entityTaskCategory = "TASK_CATEGORY"
	entityTaskGoal     = "TASK_GOAL"
	entityTaskGroup    = "TASK_GROUP"
	entityTag          = "TAG"
	entityTaskTag      = "TASK_TAG"

	operationAdd      = "ADD"
	operationUpdate   = "UPDATE"
	operationDelete   = "DELETE"
	operationComplete = "COMPLETE"
)

type Transaction struct {
	ID        string
	Entity    string
	Operation string
	Args      json.RawMessage
	At        int
}

type TaskContent struct {
	Content  string
	Note     string
	Start    int
	Reminder int
	Deadline int
	Order    int
}

func (tc TaskContent) CopyToTask(t *storage.Task) {
	if len(tc.Content) != 0 {
		t.Content = tc.Content
	}
	if len(tc.Note) != 0 {
		t.Note = tc.Note
	}
	if tc.Start != 0 {
		t.Start = ptrTime(time.Unix(int64(tc.Start), 0))
	}
	if tc.Reminder != 0 {
		t.Reminder = ptrTime(time.Unix(int64(tc.Reminder), 0))
	}
	if tc.Deadline != 0 {
		t.Deadline = ptrTime(time.Unix(int64(tc.Deadline), 0))
	}
	if tc.Order != 0 {
		t.Order = tc.Order
	}
}

type ChecklistContent struct {
	TaskID  string
	Content string
	Order   int
}

func (cc ChecklistContent) CopyToChecklist(c *storage.Checklist) {
	// Only update taskID when current taskID is not set
	if len(c.TaskID) == 0 {
		c.TaskID = cc.TaskID
	}
	if len(cc.Content) != 0 {
		c.Content = cc.Content
	}
	if cc.Order != 0 {
		c.Order = cc.Order
	}
}

func ptrTime(t time.Time) *time.Time {
	return &t
}
