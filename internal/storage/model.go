package storage

import (
	uuid "github.com/satori/go.uuid"
	"time"
	"wetodo/internal/account"
	"wetodo/internal/operation"
	"wetodo/internal/user"
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

// OauthToken ...
type OauthToken struct {
	ID           int `gorm:"primary_key"`
	UserID       int
	AccessToken  string
	ExpiresAt    time.Time
	RefreshToken string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// CopyToModel copy data from GORM model to servide model
func (u User) CopyToModel() user.User {
	return user.User{
		ID:         u.ID,
		Name:       u.Name,
		Email:      u.Email,
		Password:   u.Password,
		PictureURL: u.PictureURL,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

// CopyToModel copy data from GORM model to servide model
func (o OauthToken) CopyToModel() account.OauthToken {
	return account.OauthToken{
		ID:           o.ID,
		UserID:       o.UserID,
		AccessToken:  o.AccessToken,
		ExpiresAt:    o.ExpiresAt,
		RefreshToken: o.RefreshToken,
		CreatedAt:    o.CreatedAt,
		UpdatedAt:    o.UpdatedAt,
	}
}

type Task struct {
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
	Completed   bool
	Deleted     bool
	DeletedAt   *time.Time
	CompletedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Task) CopyToServiceModel() operation.Task {
	u, _ := uuid.FromBytes(t.ID)
	return operation.Task{
		ID:          u.String(),
		UserID:      t.UserID,
		TaskGroupID: t.TaskGroupID,
		TaskGoalID:  t.TaskGoalID,
		Content:     t.Content,
		Note:        t.Note,
		Start:       t.Start,
		Reminder:    t.Reminder,
		Deadline:    t.Deadline,
		Order:       t.Order,
		Completed:   t.Completed,
		Deleted:     t.Deleted,
		DeletedAt:   t.DeletedAt,
		CompletedAt: t.CompletedAt,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func CopyTaskFromServiceModel(t operation.Task) Task {
	uID, _ := uuidToBinary(t.ID)
	return Task{
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
		Completed:   t.Completed,
		Deleted:     t.Deleted,
		DeletedAt:   t.DeletedAt,
		CompletedAt: t.CompletedAt,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
