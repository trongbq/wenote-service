package request

import "wetodo/internal/storage"

// RegisterRequest contains request data for SignUp handler
type RegisterRequest struct {
	Name     string
	Email    string
	Password string
}

// SignInRequest contains request data for SignIn handler
type LoginRequest struct {
	Email    string
	Password string
}

// SignInRequest contains request data for SignIn handler
type RefreshOauthTokenRequest struct {
	RefreshToken string
}

// CopyToModel ...
func (r RegisterRequest) CopyToModel() storage.User {
	return storage.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}
