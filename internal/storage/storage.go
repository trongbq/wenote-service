package storage

import (
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

// GetAllUsers return all user in db
func (s *Storage) GetAllUsers() []user.User {
	users := []user.User{}
	s.db.Find(&users)
	return users
}

// GetUserByID return single user contains matched ID
func (s *Storage) GetUserByID(id int) (user.User, bool) {
	var user user.User
	if s.db.First(&user, id).RecordNotFound() {
		return user, false
	}
	return user, true
}

// CreateUser save user data into DB
func (s *Storage) CreateUser(u user.User) {
	s.db.Create(u)
}
