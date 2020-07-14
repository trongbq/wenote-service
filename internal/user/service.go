package user

// Repository provides access to user repository
type Repository interface {
	GetAllUsers() []User
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
