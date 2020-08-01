package transaction

import (
	"time"
)

const (
	MethodSet    = "SET"
	MethodDelete = "DELETE"

	EntityTask      = "TASK"
	EntityChecklist = "Checklist"
	EntityTaskGoal  = "TaskGoal"
	EntityTaskGroup = "TaskGroup"
	EntityTag       = "Tag"
	EntityTaskTag   = "TaskTag"
)

type Transaction struct {
	ID      string
	Entity  string
	Actions []Action
}

// Action records user actions on specific object
type Action struct {
	Method    string
	Argument  Argument
	StartedAt time.Time
}

type Argument struct {
	Name  string
	Value string
}
