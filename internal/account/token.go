package account

import (
	"fmt"
	"github.com/pkg/errors"
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

// ExtractUserIDFromToken returns userID from token
func ExtractUserIDFromToken(token string) (int, error) {
	claims, err := extractToken(token)
	if err != nil {
		return -1, err
	}

	if !verifyIssuer(claims["iss"].(string)) {
		return -1, errors.New("Invalid token issuer")
	}

	// Bug: userId has type float64 instead of int
	return int(claims["userId"].(float64)), nil
}

func generateToken(userID int, tokenType string) (Token, error) {
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

func verifyToken(token string) bool {
	claims, err := extractToken(token)
	if err != nil {
		return false
	}
	return verifyIssuer(claims["iss"].(string))
}

func verifyIssuer(issuer string) bool {
	return issuer == viper.GetString("token.issuer")
}

func extractToken(token string) (map[string]interface{}, error) {
	extractToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("token.secret")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := extractToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Error extract token")
	}
	if claims.Valid() != nil {
		return nil, errors.New("Invalid claim")
	}

	return claims, nil
}
