package storage

import (
	"time"
	"wenote/internal/account"
	"wenote/internal/user"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // gorm dialect
)

// Storage store db connection
type Storage struct {
	db *gorm.DB
}

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
func (s *Storage) CreateUser(u user.User) (user.User, error) {
	uStorage := User{
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.db.Save(&uStorage).Error; err != nil {
		return u, err
	}
	newUser := user.User{
		ID:         uStorage.ID,
		Name:       uStorage.Name,
		Email:      uStorage.Email,
		Password:   uStorage.Password,
		PictureURL: uStorage.PictureURL,
		CreatedAt:  uStorage.CreatedAt,
		UpdatedAt:  uStorage.UpdatedAt,
	}
	return newUser, nil
}

// GetOauthTokenByUserID ...
func (s *Storage) GetOauthTokenByUserID(userID int) (account.OauthToken, bool) {
	var auth account.OauthToken
	if s.db.Where("user_id = ?", userID).First(&auth).RecordNotFound() {
		return auth, false
	}
	return auth, true
}

// CreateOauthToken ...
func (s *Storage) CreateOauthToken(auth account.OauthToken) (account.OauthToken, error) {
	at := OauthToken{
		UserID:       auth.UserID,
		AccessToken:  auth.AccessToken,
		ExpiresAt:    auth.ExpiresAt,
		RefreshToken: auth.RefreshToken,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := s.db.Save(&at).Error; err != nil {
		return auth, err
	}
	newAuth := account.OauthToken{
		ID:           at.ID,
		UserID:       at.UserID,
		AccessToken:  at.AccessToken,
		ExpiresAt:    at.ExpiresAt,
		RefreshToken: at.RefreshToken,
		CreatedAt:    at.CreatedAt,
		UpdatedAt:    at.UpdatedAt,
	}
	return newAuth, nil
}
