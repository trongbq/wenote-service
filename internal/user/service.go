package user

import (
	"wetodo/internal/storage"
)

// Repository provides access to user repository
type Repository interface {
	GetAllUsers() []storage.User
	GetUserByID(id int) (storage.User, bool)
	CreateUser(u storage.User) (storage.User, error)
}

// Service provides user operations
type Service struct {
	r Repository
}

// NewService creates a user service
func NewService(r Repository) *Service {
	return &Service{r}
}

// GetAllUsers returns all users
func (s *Service) GetAllUsers() []storage.User {
	return s.r.GetAllUsers()
}

// GetUserByID return single user
func (s *Service) GetUserByID(id int) (storage.User, bool) {
	return s.r.GetUserByID(id)
}

// CreateUser creates user with name, email and password
func (s *Service) CreateUser(name string, email string, password string) (storage.User, error) {
	u := storage.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	return s.r.CreateUser(u)
}
