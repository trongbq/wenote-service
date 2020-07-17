package storage

import (
	"github.com/spf13/viper"
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

// NewStorage return a new MySQL storage
func NewStorage() (*Storage, error) {
	db, err := gorm.Open("mysql", viper.GetString("database.connection-url"))
	if err != nil {
		return nil, err
	}

	db.SingularTable(true)

	if env := viper.GetString("env"); env == "local" {
		db.LogMode(true)
	}

	return &Storage{db}, nil
}

// GetAllUsers return all user in db
func (s *Storage) GetAllUsers() []user.User {
	var users []user.User
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

// GetUserByEmail return single user contains matched email
func (s *Storage) GetUserByEmail(email string) (user.User, bool) {
	var user user.User
	if s.db.Where("email = ?", email).First(&user).RecordNotFound() {
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
	return uStorage.CopyToModel(), nil
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
	return at.CopyToModel(), nil
}
