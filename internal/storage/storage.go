package storage

import (
	uuid "github.com/satori/go.uuid"
	"time"
	"wetodo/internal/account"
	"wetodo/internal/operation"
	"wetodo/internal/user"

	"github.com/spf13/viper"

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

// UpdateOauthToken ...
func (s *Storage) UpdateOauthToken(auth account.OauthToken) (account.OauthToken, error) {
	at := OauthToken{
		ID:           auth.ID,
		UserID:       auth.UserID,
		AccessToken:  auth.AccessToken,
		ExpiresAt:    auth.ExpiresAt,
		RefreshToken: auth.RefreshToken,
		CreatedAt:    auth.CreatedAt,
		UpdatedAt:    time.Now(),
	}
	if err := s.db.Save(&at).Error; err != nil {
		return auth, err
	}
	return at.CopyToModel(), nil
}

// DeleteOauthTokenByUserID deletes user credentials
func (s *Storage) DeleteOauthTokenByUserID(userID int) error {
	return s.db.Where("user_id = ?", userID).Delete(OauthToken{}).Error
}

// CreateOrUpdateTask creates or updates task
func (s *Storage) CreateOrUpdateTask(t operation.Task) (operation.Task, error) {
	task := CopyTaskFromServiceModel(t)
	task.UpdatedAt = time.Now()
	if err := s.db.Save(&task).Error; err != nil {
		return task.CopyToServiceModel(), err
	}

	return task.CopyToServiceModel(), nil
}

// GetOauthTokenByUserID returns a task of user with specific ID
func (s *Storage) GetTaskByID(userID int, id string) (operation.Task, bool) {
	var task Task

	uID, err := uuidToBinary(id)
	if err != nil {
		return operation.Task{}, false
	}

	if s.db.Where("user_id = ? AND id = ?", userID, uID).First(&task).RecordNotFound() {
		return operation.Task{}, false
	}
	return task.CopyToServiceModel(), true
}

func uuidToBinary(s string) ([]byte, error) {
	u, err := uuid.FromString(s)
	if err != nil {
		return []byte{}, err
	}
	return u.MarshalBinary()
}
