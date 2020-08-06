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

type TaskCategoryContent struct {
	Name  string
	Order int
}

func (tc TaskCategoryContent) CopyToTask(t *storage.TaskCategory) {
	if len(tc.Name) != 0 {
		t.Name = tc.Name
	}
	if tc.Order != 0 {
		t.Order = tc.Order
	}
}

type TaskGoalContent struct {
	CategoryID string
	Content    string
	Note       string
	Start      int
	Reminder   int
	Deadline   int
	Order      int
}

func (tg TaskGoalContent) CopyToTask(t *storage.TaskGoal) {
	if len(tg.CategoryID) != 0 {
		t.CategoryID = tg.CategoryID
	}
	if len(tg.Content) != 0 {
		t.Content = tg.Content
	}
	if len(tg.Note) != 0 {
		t.Note = tg.Note
	}
	if tg.Start != 0 {
		t.Start = ptrTime(time.Unix(int64(tg.Start), 0))
	}
	if tg.Reminder != 0 {
		t.Reminder = ptrTime(time.Unix(int64(tg.Reminder), 0))
	}
	if tg.Deadline != 0 {
		t.Deadline = ptrTime(time.Unix(int64(tg.Deadline), 0))
	}
	if tg.Order != 0 {
		t.Order = tg.Order
	}
}

type TaskGroupContent struct {
	TaskGoalID string
	Content    string
	Order      int
}

func (tg TaskGroupContent) CopyToTask(t *storage.TaskGroup) {
	if len(tg.TaskGoalID) != 0 {
		t.TaskGoalID = tg.TaskGoalID
	}
	if len(tg.Content) != 0 {
		t.Content = tg.Content
	}
	if tg.Order != 0 {
		t.Order = tg.Order
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

type TagContent struct {
	Name string
}

type TaskTagContent struct {
	TaskID string
	TagID  string
}

func ptrTime(t time.Time) *time.Time {
	return &t
}
