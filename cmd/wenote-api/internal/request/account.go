package request

import (
	"wenote/internal/user"
)

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

// CopyToModel ...
func (r RegisterRequest) CopyToModel() user.User {
	return user.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}
