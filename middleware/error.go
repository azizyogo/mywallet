package middleware

import (
	"log"
	"mywallet/apperror"
	"mywallet/pkg/httputil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ErrorHandler middleware handles panics and converts them to proper HTTP responses
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				httputil.SendError(c, http.StatusInternalServerError, "Internal server error", nil)
			}
		}()
		c.Next()
	}
}

// ValidationErrorResponse formats validation errors
func ValidationErrorResponse(err error) interface{} {
	var errors []map[string]string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errors = append(errors, map[string]string{
				"field":   e.Field(),
				"message": getValidationMessage(e),
			})
		}
	}

	return errors
}

func getValidationMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short (minimum: " + e.Param() + ")"
	case "max":
		return "Value is too long (maximum: " + e.Param() + ")"
	case "gt":
		return "Value must be greater than " + e.Param()
	case "gte":
		return "Value must be greater than or equal to " + e.Param()
	default:
		return "Invalid value"
	}
}

// HandleAppError handles application-specific errors
func HandleAppError(c *gin.Context, err error) {
	if appErr, ok := err.(*apperror.AppError); ok {
		httputil.SendError(c, appErr.StatusCode, appErr.Message, nil)
		return
	}
	httputil.SendError(c, http.StatusInternalServerError, "Internal server error", nil)
}
