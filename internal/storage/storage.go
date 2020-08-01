package storage

import (
	"time"

	uuid "github.com/satori/go.uuid"

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
func (s *Storage) GetAllUsers() []User {
	var users []User
	s.db.Find(&users)
	return users
}

// GetUserByID return single user contains matched ID
func (s *Storage) GetUserByID(id int) (User, bool) {
	var user User
	if s.db.First(&user, id).RecordNotFound() {
		return user, false
	}
	return user, true
}

// GetUserByEmail return single user contains matched email
func (s *Storage) GetUserByEmail(email string) (User, bool) {
	var user User
	if s.db.Where("email = ?", email).First(&user).RecordNotFound() {
		return user, false
	}
	return user, true
}

// CreateUser save user data into DB
func (s *Storage) CreateUser(u User) (User, error) {
	user := User{
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.db.Save(&user).Error; err != nil {
		return u, err
	}
	return user, nil
}

// GetOauthTokenByUserID ...
func (s *Storage) GetOauthTokenByUserID(userID int) (OauthToken, bool) {
	var auth OauthToken
	if s.db.Where("user_id = ?", userID).First(&auth).RecordNotFound() {
		return auth, false
	}
	return auth, true
}

// CreateOauthToken ...
func (s *Storage) CreateOauthToken(auth OauthToken) (OauthToken, error) {
	authToken := OauthToken{
		UserID:       auth.UserID,
		AccessToken:  auth.AccessToken,
		ExpiresAt:    auth.ExpiresAt,
		RefreshToken: auth.RefreshToken,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := s.db.Save(&authToken).Error; err != nil {
		return auth, err
	}
	return authToken, nil
}

// UpdateOauthToken ...
func (s *Storage) UpdateOauthToken(auth OauthToken) (OauthToken, error) {
	authToken := OauthToken{
		ID:           auth.ID,
		UserID:       auth.UserID,
		AccessToken:  auth.AccessToken,
		ExpiresAt:    auth.ExpiresAt,
		RefreshToken: auth.RefreshToken,
		CreatedAt:    auth.CreatedAt,
		UpdatedAt:    time.Now(),
	}
	if err := s.db.Save(&authToken).Error; err != nil {
		return auth, err
	}
	return authToken, nil
}

// DeleteOauthTokenByUserID deletes user credentials
func (s *Storage) DeleteOauthTokenByUserID(userID int) error {
	return s.db.Where("user_id = ?", userID).Delete(OauthToken{}).Error
}

// CreateOrUpdateTask creates or updates task
func (s *Storage) CreateOrUpdateTask(t Task) (Task, error) {
	task := t.CopyToInternalModel()
	task.UpdatedAt = time.Now()

	if err := s.db.Save(&task).Error; err != nil {
		return t, err
	}

	return task.CopyToRepModel(), nil
}

// GetOauthTokenByUserID returns a task of user with specific ID
func (s *Storage) GetTaskByID(userID int, id string) (Task, bool) {
	var task TaskInternal

	uID, err := uuidToBinary(id)
	if err != nil {
		return Task{}, false
	}

	if s.db.Where("user_id = ? AND id = ?", userID, uID).First(&task).RecordNotFound() {
		return Task{}, false
	}
	return task.CopyToRepModel(), true
}

func uuidToBinary(s string) ([]byte, error) {
	u, err := uuid.FromString(s)
	if err != nil {
		return []byte{}, err
	}
	return u.MarshalBinary()
}
