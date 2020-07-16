package account

import "errors"

// ErrFailedGenerateToken indicates token generation is imcomplete
var ErrFailedGenerateToken = errors.New("Can not generate account token")

// ErrDupplicateEmail indicates email used for registering is used in the system
var ErrDupplicateEmail = errors.New("Email is used")
