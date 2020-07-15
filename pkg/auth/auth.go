package auth

import "golang.org/x/crypto/bcrypt"

// HashAndSalt creates hashed string
func HashAndSalt(s string) (string, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(p), nil
}

// CompareHashAndPassword compares hash string and plain password
func CompareHashAndPassword(h string, p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	if err != nil {
		return false
	}
	return true
}
