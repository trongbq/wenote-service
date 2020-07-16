package account

import "time"

// OauthToken ...
type OauthToken struct {
	ID           int
	UserID       int
	AccessToken  string
	ExpiresAt    time.Time
	RefreshToken string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
