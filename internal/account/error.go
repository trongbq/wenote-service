package account

import "errors"

// ErrFailedGenerateToken indicates token generation is imcomplete
var ErrFailedGenerateToken = errors.New("can not generate account token")

// ErrDupplicateEmail indicates email used for registering is used in the system
var ErrDuplicateEmail = errors.New("email is used")

// ErrUserNotFound indicates user does not exist in the system
var ErrUserNotFound = errors.New("user is not found")

// ErrUserNotFound indicates user does not exist in the system
var ErrInvalidPassword = errors.New("password is invalid")

// ErrUserNotFound indicates user does not exist in the system
var ErrInvalidRefreshToken = errors.New("refresh token is invalid")
