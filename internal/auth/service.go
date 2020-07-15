package auth

import (
	"wenote/internal/user"

	"golang.org/x/crypto/bcrypt"
)

// Service provides user operations
type Service struct {
	u *user.Service
}

// NewService creates a user service
func NewService(u *user.Service) *Service {
	return &Service{u}
}

// SignUp sign user up
func (s *Service) SignUp(name string, email string, password string) (*user.User, error) {
	// Hash password
	p, err := hashAndSalt(password)
	if err != nil {
		return nil, err
	}
	// Create user
	u, err := s.u.CreateUser(name, email, p)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func hashAndSalt(s string) (string, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(p), nil
}

func compareHashAndPassword(h string, p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	if err != nil {
		return false
	}
	return true
}
