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
	TaskGroupID string
	TaskGoalID  string
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

type TaskGroup struct {
	ID         string
	UserID     int
	TaskGoalID string
	Name       string
	Order      int
	DeletedAt  *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type TaskGoal struct {
	ID          string
	UserID      int
	CategoryID  string
	Name        string
	Note        string
	Start       *time.Time
	Reminder    *time.Time
	Deadline    *time.Time
	Order       int
	CompletedAt *time.Time
	DeletedAt   *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TaskCategory struct {
	ID        string
	UserID    int
	Name      string
	Order     int
	DeletedAt *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Checklist struct {
	ID          string
	TaskID      string
	Content     string
	Order       int
	CompletedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Tag struct {
	ID        string
	UserID    int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TaskInternal struct {
	ID          []byte
	UserID      int
	TaskGroupID []byte
	TaskGoalID  []byte
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

func (t Task) CopyToInternalModel() TaskInternal {
	id, _ := uuidToBinary(t.ID)
	tgrID, _ := uuidToBinary(t.TaskGroupID)
	utD, _ := uuidToBinary(t.TaskGoalID)
	return TaskInternal{
		ID:          id,
		UserID:      t.UserID,
		TaskGroupID: tgrID,
		TaskGoalID:  utD,
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
	id, _ := uuid.FromBytes(t.ID)
	tgrID, _ := uuid.FromBytes(t.TaskGroupID)
	tgID, _ := uuid.FromBytes(t.TaskGoalID)
	return Task{
		ID:          id.String(),
		UserID:      t.UserID,
		TaskGroupID: tgrID.String(),
		TaskGoalID:  tgID.String(),
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

type ChecklistInternal struct {
	ID          []byte
	TaskID      []byte
	Content     string
	Order       int
	CompletedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (ChecklistInternal) TableName() string {
	return "checklist"
}

func (c Checklist) CopyToInternalModel() ChecklistInternal {
	id, _ := uuidToBinary(c.ID)
	tID, _ := uuidToBinary(c.TaskID)
	return ChecklistInternal{
		ID:          id,
		TaskID:      tID,
		Content:     c.Content,
		Order:       c.Order,
		CompletedAt: c.CompletedAt,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

func (c ChecklistInternal) CopyToRepModel() Checklist {
	id, _ := uuid.FromBytes(c.ID)
	tID, _ := uuid.FromBytes(c.TaskID)
	return Checklist{
		ID:        id.String(),
		TaskID:    tID.String(),
		Content:   c.Content,
		Order:     c.Order,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

type TagInternal struct {
	ID        []byte
	UserID    int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (TagInternal) TableName() string {
	return "tag"
}

func (t *Tag) CopyToInternalModel() TagInternal {
	id, _ := uuidToBinary(t.ID)
	return TagInternal{
		ID:        id,
		UserID:    t.UserID,
		Name:      t.Name,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func (t TagInternal) CopyToRepModel() Tag {
	id, _ := uuid.FromBytes(t.ID)
	return Tag{
		ID:        id.String(),
		UserID:    t.UserID,
		Name:      t.Name,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

type TaskCategoryInternal struct {
	ID        []byte
	UserID    int
	Name      string
	Order     int
	DeletedAt *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (TaskCategoryInternal) TableName() string {
	return "task_category"
}

func (t *TaskCategory) CopyToInternalModel() TaskCategoryInternal {
	id, _ := uuidToBinary(t.ID)
	return TaskCategoryInternal{
		ID:        id,
		UserID:    t.UserID,
		Name:      t.Name,
		Order:     t.Order,
		DeletedAt: t.DeletedAt,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func (t TaskCategoryInternal) CopyToRepModel() TaskCategory {
	id, _ := uuid.FromBytes(t.ID)
	return TaskCategory{
		ID:        id.String(),
		UserID:    t.UserID,
		Name:      t.Name,
		Order:     t.Order,
		DeletedAt: t.DeletedAt,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

type TaskGoalInternal struct {
	ID          []byte
	UserID      int
	CategoryID  []byte
	Name        string
	Note        string
	Start       *time.Time
	Reminder    *time.Time
	Deadline    *time.Time
	Order       int
	CompletedAt *time.Time
	DeletedAt   *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (TaskGoalInternal) TableName() string {
	return "task_goal"
}

func (t *TaskGoal) CopyToInternalModel() TaskGoalInternal {
	id, _ := uuidToBinary(t.ID)
	cID, _ := uuidToBinary(t.CategoryID)
	return TaskGoalInternal{
		ID:          id,
		UserID:      t.UserID,
		CategoryID:  cID,
		Name:        t.Name,
		Note:        t.Note,
		Start:       t.Start,
		Reminder:    t.Reminder,
		Deadline:    t.Deadline,
		Order:       t.Order,
		CompletedAt: t.CompletedAt,
		DeletedAt:   t.DeletedAt,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func (t TaskGoalInternal) CopyToRepModel() TaskGoal {
	id, _ := uuid.FromBytes(t.ID)
	cID, _ := uuid.FromBytes(t.CategoryID)
	return TaskGoal{
		ID:          id.String(),
		UserID:      t.UserID,
		CategoryID:  cID.String(),
		Name:        t.Name,
		Note:        t.Note,
		Start:       t.Start,
		Reminder:    t.Reminder,
		Deadline:    t.Deadline,
		Order:       t.Order,
		CompletedAt: t.CompletedAt,
		DeletedAt:   t.DeletedAt,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

type TaskGroupInternal struct {
	ID         []byte
	UserID     int
	TaskGoalID []byte
	Name       string
	Order      int
	DeletedAt  *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (TaskGroupInternal) TableName() string {
	return "task_group"
}

func (t *TaskGroup) CopyToInternalModel() TaskGroupInternal {
	id, _ := uuidToBinary(t.ID)
	tID, _ := uuidToBinary(t.TaskGoalID)
	return TaskGroupInternal{
		ID:         id,
		UserID:     t.UserID,
		TaskGoalID: tID,
		Name:       t.Name,
		Order:      t.Order,
		DeletedAt:  t.DeletedAt,
		CreatedAt:  t.CreatedAt,
		UpdatedAt:  t.UpdatedAt,
	}
}

func (t TaskGroupInternal) CopyToRepModel() TaskGroup {
	id, _ := uuid.FromBytes(t.ID)
	tID, _ := uuid.FromBytes(t.TaskGoalID)
	return TaskGroup{
		ID:         id.String(),
		UserID:     t.UserID,
		TaskGoalID: tID.String(),
		Name:       t.Name,
		Order:      t.Order,
		DeletedAt:  t.DeletedAt,
		CreatedAt:  t.CreatedAt,
		UpdatedAt:  t.UpdatedAt,
	}
}

type TaskTagInternal struct {
	tagID     []byte
	taskID    []byte
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
