package operation

import (
	"time"
	"wetodo/internal/storage"
)

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
	Content   *OperationContent // For add or update task
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

func (o OperationContent) CopyToTask() storage.Task {
	return storage.Task{
		Content:  o.Content,
		Note:     o.Note,
		Start:    o.Start,
		Reminder: o.Reminder,
		Deadline: o.Deadline,
		Order:    o.Order,
	}
}

func (o OperationContent) UpdateToTask(task storage.Task) storage.Task {
	task.Content = o.Content
	task.Note = o.Note
	task.Start = o.Start
	task.Reminder = o.Reminder
	task.Deadline = o.Deadline
	task.Order = o.Order
	return task
}
