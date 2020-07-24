package response

import (
	"time"
	"wetodo/internal/account"
)

// AccountTokenResponse ...
type AccountTokenResponse struct {
	UserID       int       `json:"userId"`
	AccessToken  string    `json:"accessToken"`
	ExpiresAt    time.Time `json:"expiresAt"`
	RefreshToken string    `json:"refreshToken"`
}

// CopyToAccountRegisterResponse ...
func CopyToAccountRegisterResponse(u account.OauthToken) AccountTokenResponse {
	return AccountTokenResponse{
		UserID:       u.UserID,
		AccessToken:  u.AccessToken,
		ExpiresAt:    u.ExpiresAt,
		RefreshToken: u.RefreshToken,
	}
}
