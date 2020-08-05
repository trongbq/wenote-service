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

// GetTaskByID returns a task of user with specific ID
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

// CreateOrUpdateChecklist creates or updates checklist
func (s *Storage) CreateOrUpdateChecklist(cl Checklist) (Checklist, error) {
	checklist := cl.CopyToInternalModel()
	checklist.UpdatedAt = time.Now()

	if err := s.db.Save(&checklist).Error; err != nil {
		return cl, err
	}

	return checklist.CopyToRepModel(), nil
}

// GetChecklistByID returns a checklist of user with specific ID
func (s *Storage) GetChecklistByID(id string) (Checklist, bool) {
	var cl ChecklistInternal

	uID, err := uuidToBinary(id)
	if err != nil {
		return Checklist{}, false
	}

	if s.db.Where("id = ?", uID).First(&cl).RecordNotFound() {
		return Checklist{}, false
	}
	return cl.CopyToRepModel(), true
}

// DeleteChecklistByID delete a checklist by its ID
func (s *Storage) DeleteChecklistByID(id string) {
	cl := ChecklistInternal{ID: uuidToBinaryShort(id)}

	s.db.Delete(&cl)
}

// CreateOrUpdateTaskCategory creates or updates task category
func (s *Storage) CreateOrUpdateTaskCategory(tc TaskCategory) (TaskCategory, error) {
	taskCat := tc.CopyToInternalModel()
	taskCat.UpdatedAt = time.Now()

	if err := s.db.Save(&taskCat).Error; err != nil {
		return tc, err
	}

	return taskCat.CopyToRepModel(), nil
}

// GetTaskCategoryByID returns a task category of user with specific ID
func (s *Storage) GetTaskCategoryByID(userID int, id string) (TaskCategory, bool) {
	var tc TaskCategoryInternal

	uID, err := uuidToBinary(id)
	if err != nil {
		return TaskCategory{}, false
	}

	if s.db.Where("user_id = ? AND id = ?", userID, uID).First(&tc).RecordNotFound() {
		return TaskCategory{}, false
	}
	return tc.CopyToRepModel(), true
}

// CreateOrUpdateTaskGoal creates or updates task category
func (s *Storage) CreateOrUpdateTaskGoal(tg TaskGoal) (TaskGoal, error) {
	taskGoal := tg.CopyToInternalModel()
	taskGoal.UpdatedAt = time.Now()

	if err := s.db.Save(&taskGoal).Error; err != nil {
		return tg, err
	}

	return taskGoal.CopyToRepModel(), nil
}

// GetTaskGoalByID returns a task goal of user with specific ID
func (s *Storage) GetTaskGoalByID(userID int, id string) (TaskGoal, bool) {
	var tg TaskGoalInternal

	uID, err := uuidToBinary(id)
	if err != nil {
		return TaskGoal{}, false
	}

	if s.db.Where("user_id = ? AND id = ?", userID, uID).First(&tg).RecordNotFound() {
		return TaskGoal{}, false
	}
	return tg.CopyToRepModel(), true
}

// CreateOrUpdateTaskGroup creates or updates task group
func (s *Storage) CreateOrUpdateTaskGroup(tg TaskGroup) (TaskGroup, error) {
	taskGroup := tg.CopyToInternalModel()
	taskGroup.UpdatedAt = time.Now()

	if err := s.db.Save(&taskGroup).Error; err != nil {
		return tg, err
	}

	return taskGroup.CopyToRepModel(), nil
}

// GetTaskGroupByID returns a task category of user with specific ID
func (s *Storage) GetTaskGroupByID(userID int, id string) (TaskGroup, bool) {
	var tg TaskGroupInternal

	uID, err := uuidToBinary(id)
	if err != nil {
		return TaskGroup{}, false
	}

	if s.db.Where("user_id = ? AND id = ?", userID, uID).First(&tg).RecordNotFound() {
		return TaskGroup{}, false
	}
	return tg.CopyToRepModel(), true
}

func uuidToBinary(s string) ([]byte, error) {
	u, err := uuid.FromString(s)
	if err != nil {
		return []byte{}, err
	}
	return u.MarshalBinary()
}

func uuidToBinaryShort(s string) []byte {
	u, err := uuid.FromString(s)
	if err != nil {
		return []byte{}
	}
	v, _ := u.MarshalBinary()
	return v
}
