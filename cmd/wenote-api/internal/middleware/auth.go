package middleware

import (
	"net/http"
	"strings"
	"time"
	"wenote/cmd/wenote-api/internal/error"
	"wenote/internal/account"

	"github.com/gin-gonic/gin"
)

const BearerToken = "Bearer"

// AuthenticationMiddleware authenticates request and set value to context
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if len(authHeader) == 0 || !strings.HasPrefix(authHeader, BearerToken) || len(strings.Fields(authHeader)) != 2 {
			c.JSON(http.StatusUnauthorized, error.SimpleUnauthorizedResponse("Invalid authorization header"))
			c.Abort()
			return
		}

		authToken := strings.Fields(authHeader)[1]
		userID, err := account.ExtractUserIDFromToken(authToken)
		if err != nil {
			if err.Error() == "Token is expired" {
				c.JSON(http.StatusUnauthorized, error.Error{
					Code:      error.ErrorCodeTokenExpired,
					Message:   err.Error(),
					Timestamp: time.Now(),
				})
			} else {
				c.JSON(http.StatusUnauthorized, error.SimpleUnauthorizedResponse(err.Error()))
			}
			c.Abort()
			return
		}

		// TODO: Verify that user has this authentication token
		// TODO: use cache storage for fast retrieving

		c.Set("UserID", userID)
		c.Next()
	}
}
