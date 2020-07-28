package operation

import "time"

const (
	AddTask      = "ADD_TASK"
	RemoveTask   = "REMOVE_TASK"
	UpdateTask   = "UPDATE_TASK"
	CompleteTask = "COMPLETE_TASK"
)

// Operation records user actions
type Operation struct {
	ID        string // Target task ID
	Type      string
	Content   *OperationContent
	StartedAt time.Time
}

type OperationContent struct {
	Content  string
	Note     string
	Start    *time.Time
	Reminder *time.Time
	Deadline *time.Time
	Order    int
}

type Task struct {
	ID          string
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
	Deleted     bool
	DeletedAt   *time.Time
	CompletedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (o OperationContent) CopyToTask() Task {
	return Task{
		Content:  o.Content,
		Note:     o.Note,
		Start:    o.Start,
		Reminder: o.Reminder,
		Deadline: o.Deadline,
		Order:    o.Order,
	}
}

func (o OperationContent) UpdateToTask(task Task) Task {
	task.Content = o.Content
	task.Note = o.Note
	task.Start = o.Start
	task.Reminder = o.Reminder
	task.Deadline = o.Deadline
	task.Order = o.Order
	return task
}
