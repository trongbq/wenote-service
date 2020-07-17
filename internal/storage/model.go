package storage

import (
	"time"
	"wenote/internal/account"
	"wenote/internal/user"
)

// User type in  GORM
type User struct {
	ID         int `gorm:"primary_key"`
	Name       string
	Email      string
	PictureURL string
	Password   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// OauthToken ...
type OauthToken struct {
	ID           int `gorm:"primary_key"`
	UserID       int
	AccessToken  string
	ExpiresAt    time.Time
	RefreshToken string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// CopyToModel copy data from GORM model to servide model
func (u User) CopyToModel() user.User {
	return user.User{
		ID:         u.ID,
		Name:       u.Name,
		Email:      u.Email,
		Password:   u.Password,
		PictureURL: u.PictureURL,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

// CopyToModel copy data from GORM model to servide model
func (o OauthToken) CopyToModel() account.OauthToken {
	return account.OauthToken{
		ID:           o.ID,
		UserID:       o.UserID,
		AccessToken:  o.AccessToken,
		ExpiresAt:    o.ExpiresAt,
		RefreshToken: o.RefreshToken,
		CreatedAt:    o.CreatedAt,
		UpdatedAt:    o.UpdatedAt,
	}
}
