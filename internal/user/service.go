package user

import "golang.org/x/crypto/bcrypt"

// Repository provides access to user repository
type Repository interface {
	GetAllUsers() []User
	GetUserByID(id int) (User, bool)
	CreateUser(u User)
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
func (s *Service) GetAllUsers() []User {
	return s.r.GetAllUsers()
}

// GetUserByID return single user
func (s *Service) GetUserByID(id int) (User, bool) {
	return s.r.GetUserByID(id)
}

// CreateUser creates user with name, email and password
func (s *Service) CreateUser(name string, email string, password string) (*User, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	u := User{
		Name:     name,
		Email:    email,
		Password: string(p),
	}
	s.r.CreateUser(u)
	return &u, nil
}
