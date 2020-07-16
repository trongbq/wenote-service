package account

import (
	"fmt"
	"time"
	"wenote/internal/user"

	"golang.org/x/crypto/bcrypt"
)

// Repository ...
type Repository interface {
	GetOauthTokenByUserID(userID int) (OauthToken, bool)
	CreateOauthToken(auth OauthToken) (OauthToken, error)
}

// Service provides user operations
type Service struct {
	u *user.Service
	r Repository
}

// NewService creates a user service
func NewService(u *user.Service, r Repository) *Service {
	return &Service{u, r}
}

// Register user
func (s *Service) Register(u *user.User) (OauthToken, error) {
	var auth OauthToken
	// Hash password
	p, err := hashAndSalt(u.Password)
	if err != nil {
		return auth, err
	}
	// Create user
	newUser, err := s.u.CreateUser(u.Name, u.Email, p)
	if err != nil {
		return auth, err
	}

	refToken, err := GenerateToken(newUser.ID, TokenTypeRefresh)
	if err != nil {
		fmt.Println(err)
		return auth, ErrFailedGenerateToken
	}
	aToken, err := GenerateToken(newUser.ID, TokenTypeAccess)
	if err != nil {
		fmt.Println(err)
		return auth, ErrFailedGenerateToken
	}
	auth.UserID = newUser.ID
	auth.AccessToken = aToken.Value
	auth.ExpiresAt = time.Unix(aToken.ExpiresAt, 0)
	auth.RefreshToken = refToken.Value

	auth, err = s.r.CreateOauthToken(auth)

	return auth, nil
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
