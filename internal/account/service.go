package account

import (
	"fmt"
	"time"
	"wetodo/internal/storage"

	"github.com/sirupsen/logrus"

	"golang.org/x/crypto/bcrypt"
)

// Repository ...
type Repository interface {
	GetUserByEmail(email string) (storage.User, bool)
	CreateUser(u storage.User) (storage.User, error)
	GetOauthTokenByUserID(userID int) (storage.OauthToken, bool)
	CreateOauthToken(auth storage.OauthToken) (storage.OauthToken, error)
	UpdateOauthToken(auth storage.OauthToken) (storage.OauthToken, error)
	DeleteOauthTokenByUserID(userID int) error
}

// Service provides user operations
type Service struct {
	r Repository
}

// NewService creates a user service
func NewService(r Repository) *Service {
	return &Service{r}
}

// Register user
func (s *Service) Register(u storage.User) (storage.OauthToken, error) {
	password, err := hashAndSalt(u.Password)
	if err != nil {
		return storage.OauthToken{}, err
	}

	u.Password = password
	newUser, err := s.r.CreateUser(u)
	if err != nil {
		return storage.OauthToken{}, err
	}

	auth, err := generateUserTokens(newUser.ID)
	if err != nil {
		return storage.OauthToken{}, ErrFailedGenerateToken
	}
	return s.r.CreateOauthToken(auth)
}

// Login ...
func (s *Service) Login(email string, password string) (storage.OauthToken, error) {
	u, ok := s.r.GetUserByEmail(email)
	if !ok {
		return storage.OauthToken{}, ErrUserNotFound
	}

	if matched := compareHashAndPassword(u.Password, password); !matched {
		return storage.OauthToken{}, ErrInvalidPassword
	}

	token, ok := s.r.GetOauthTokenByUserID(u.ID)
	if !ok {
		// Generate a new one
		auth, err := generateUserTokens(u.ID)
		if err != nil {
			return storage.OauthToken{}, ErrFailedGenerateToken
		}
		return s.r.CreateOauthToken(auth)
	}
	return token, nil
}

// RefreshAccessToken validates refresh token and return new access token
func (s *Service) RefreshAccessToken(refreshToken string) (storage.OauthToken, error) {
	if !verifyToken(refreshToken) {
		return storage.OauthToken{}, ErrInvalidRefreshToken
	}

	userID, err := ExtractUserIDFromToken(refreshToken)
	if err != nil {
		return storage.OauthToken{}, ErrInvalidRefreshToken
	}

	auth, ok := s.r.GetOauthTokenByUserID(userID)
	if !ok {
		return storage.OauthToken{}, ErrInvalidRefreshToken
	}

	// Generate a new access token
	accessToken, err := generateToken(userID, TokenTypeAccess)
	if err != nil {
		fmt.Println(err)
		return storage.OauthToken{}, ErrFailedGenerateToken
	}
	auth.AccessToken = accessToken.Value
	auth.ExpiresAt = time.Unix(accessToken.ExpiresAt, 0)

	return s.r.UpdateOauthToken(auth)
}

// Logout remove user's credentials from DB
func (s *Service) Logout(userID int) {
	err := s.r.DeleteOauthTokenByUserID(userID)
	if err != nil {
		logrus.Errorf("DeleteOauthTokenByUserID: %v", err)
	}
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

// Generate access token and refresh token for new user
func generateUserTokens(userID int) (storage.OauthToken, error) {
	refreshToken, err := generateToken(userID, TokenTypeRefresh)
	if err != nil {
		fmt.Println(err)
		return storage.OauthToken{}, ErrFailedGenerateToken
	}
	accessToken, err := generateToken(userID, TokenTypeAccess)
	if err != nil {
		fmt.Println(err)
		return storage.OauthToken{}, ErrFailedGenerateToken
	}
	auth := storage.OauthToken{
		UserID:       userID,
		AccessToken:  accessToken.Value,
		ExpiresAt:    time.Unix(accessToken.ExpiresAt, 0),
		RefreshToken: refreshToken.Value,
	}
	return auth, nil
}
