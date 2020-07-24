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
	Type    string
	Content string
}

type AddTaskOperation struct {
	Order   int
	Content string
}

type RemoveTaskOperation struct {
	ID int
}

type UpdateTaskOperation struct {
	ID       int
	Content  string
	Note     string
	Start    string
	Reminder time.Time
	Deadline time.Time
	Order    int
}

type CompleteTaskOperation struct {
	ID int
}
