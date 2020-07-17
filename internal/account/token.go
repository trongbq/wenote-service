package account

import (
	"fmt"
	"github.com/spf13/viper"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	TokenTypeAccess  = "ACCESS_TOKEN"
	TokenTypeRefresh = "REFRESH_TOKEN"
)

// JWTClaims ...
type JWTClaims struct {
	jwt.StandardClaims
	UserID int `json:"userId"`
}

// Token ..
type Token struct {
	Value     string
	ExpiresAt int64
}

// GenerateToken generates token respect to `type` param
func GenerateToken(userID int, tokenType string) (Token, error) {
	var expAt int64
	if tokenType == TokenTypeAccess {
		expDuration := time.Duration(viper.GetInt64("token.expire")) * time.Minute
		expAt = time.Now().Add(expDuration).Unix()
	}
	claims := JWTClaims{
		jwt.StandardClaims{
			ExpiresAt: expAt,
			Issuer:    viper.GetString("token.issuer"),
		},
		userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(viper.GetString("token.secret")))
	if err != nil {
		return Token{}, fmt.Errorf("error JWT sign string: %v", err)
	}
	t := Token{
		Value:     signed,
		ExpiresAt: expAt,
	}
	return t, nil
}
