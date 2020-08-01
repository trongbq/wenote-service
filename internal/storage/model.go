package storage

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// User type in  GORM
type User struct {
	ID         int `gorm:"primary_key"`
	Name       string
	Email      string
	PictureURL string
	Password   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// OauthToken represent user authentication token
type OauthToken struct {
	ID           int `gorm:"primary_key"`
	UserID       int
	AccessToken  string
	ExpiresAt    time.Time
	RefreshToken string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Task is a task
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
	DeletedAt   *time.Time
	CompletedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TaskInternal struct {
	ID          []byte
	UserID      int
	TaskGroupID int
	TaskGoalID  int
	Content     string
	Note        string
	Start       *time.Time
	Reminder    *time.Time
	Deadline    *time.Time
	Order       int
	DeletedAt   *time.Time
	CompletedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (TaskInternal) TableName() string {
	return "task"
}

func (t *Task) CopyToInternalModel() TaskInternal {
	uID, _ := uuidToBinary(t.ID)
	return TaskInternal{
		ID:          uID,
		UserID:      t.UserID,
		TaskGroupID: t.TaskGroupID,
		TaskGoalID:  t.TaskGoalID,
		Content:     t.Content,
		Note:        t.Note,
		Start:       t.Start,
		Reminder:    t.Reminder,
		Deadline:    t.Deadline,
		Order:       t.Order,
		DeletedAt:   t.DeletedAt,
		CompletedAt: t.CompletedAt,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func (t TaskInternal) CopyToRepModel() Task {
	uID, _ := uuid.FromBytes(t.ID)
	return Task{
		ID:          uID.String(),
		UserID:      t.UserID,
		TaskGroupID: t.TaskGroupID,
		TaskGoalID:  t.TaskGoalID,
		Content:     t.Content,
		Note:        t.Note,
		Start:       t.Start,
		Reminder:    t.Reminder,
		Deadline:    t.Deadline,
		Order:       t.Order,
		DeletedAt:   t.DeletedAt,
		CompletedAt: t.CompletedAt,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
