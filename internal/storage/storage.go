package storage

import (
	"time"
	"wenote/internal/user"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // gorm dialect
)

// Storage store db connection
type Storage struct {
	db *gorm.DB
}

// NewStorage return a new MySQL storage
func NewStorage() (*Storage, error) {
	db, err := gorm.Open("mysql", "root:@/wenote?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}

	db.SingularTable(true)
	db.LogMode(true)

	return &Storage{db}, nil
}

// User models for gorm
type User struct {
	ID         int
	Name       string
	Email      string
	PictureURL string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// GetAllUsers return all user in db
func (s *Storage) GetAllUsers() []user.User {
	users := []user.User{}
	s.db.Find(&users)
	return users
}
