package operation

import "time"

const (
	AddTask      = "ADD_TASK"
	RemoveTask   = "REMOVE_TASK"
	UpdateTask   = "UPDATE_TASK"
	CompleteTask = "COMPLETE_TASK"

	TaskStatusActive  = "active"
	TaskSTatusDeleted = "deleted"
)

// Operation records user actions
type Operation struct {
	Type    string
	Content string
}

type AddTaskOperation struct {
	Content  string
	Note     string
	Start    *time.Time
	Reminder *time.Time
	Deadline *time.Time
	Order    int
}

type RemoveTaskOperation struct {
	ID int
}

type UpdateTaskOperation struct {
	ID       int
	Content  string
	Note     string
	Start    *time.Time
	Reminder *time.Time
	Deadline *time.Time
	Order    int
}

type CompleteTaskOperation struct {
	ID int
}

// Task ...
type Task struct {
	ID          int
	UserID      int
	TaskGroupID int
	TaskGoalID  int
	Content     string
	Note        string
	Start       *time.Time
	Reminder    *time.Time
	Deadline    *time.Time
	Order       int
	Completed   bool
	Status      string
	CompletedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// CopyToTask ...
func (o AddTaskOperation) CopyToTask() Task {
	return Task{
		Content:  o.Content,
		Note:     o.Note,
		Start:    o.Start,
		Reminder: o.Reminder,
		Deadline: o.Deadline,
		Order:    o.Order,
		Status:   TaskStatusActive,
	}
}
