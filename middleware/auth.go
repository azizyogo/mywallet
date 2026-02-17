package middleware

import (
	"mywallet/apperror"
	"mywallet/shared/utils/auth"
	"mywallet/shared/utils/httpresponse"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer "
	UserIDKey           = "user_id"
	UserEmailKey        = "user_email"
)

type AuthValidator interface {
	ValidateToken(tokenString string) (*auth.JWTClaims, error)
}

// AuthMiddleware validates JWT token and sets user context
func AuthMiddleware(validator AuthValidator) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" {
			httpresponse.SendError(c, apperror.ErrUnauthorized.StatusCode, apperror.ErrUnauthorized.Message, nil)
			c.Abort()
			return
		}

		// Check Bearer prefix
		if !strings.HasPrefix(authHeader, BearerPrefix) {
			httpresponse.SendError(c, apperror.ErrUnauthorized.StatusCode, "Invalid authorization header format", nil)
			c.Abort()
			return
		}

		// Extract token
		tokenString := strings.TrimPrefix(authHeader, BearerPrefix)

		// Validate token
		claims, err := validator.ValidateToken(tokenString)
		if err != nil {
			httpresponse.SendError(c, apperror.ErrUnauthorized.StatusCode, apperror.ErrUnauthorized.Message, nil)
			c.Abort()
			return
		}

		// Set user context
		c.Set(UserIDKey, claims.UserID)
		c.Set(UserEmailKey, claims.Email)

		c.Next()
	}
}

// GetUserID retrieves user ID from context
func GetUserID(c *gin.Context) (uint, bool) {
	userID, exists := c.Get(UserIDKey)
	if !exists {
		return 0, false
	}
	id, ok := userID.(uint)
	return id, ok
}

// GetUserEmail retrieves user email from context
func GetUserEmail(c *gin.Context) (string, bool) {
	email, exists := c.Get(UserEmailKey)
	if !exists {
		return "", false
	}
	emailStr, ok := email.(string)
	return emailStr, ok
}
