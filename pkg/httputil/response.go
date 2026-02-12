package httputil

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Meta    any    `json:"meta,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

// SendSuccess sends a successful response without metadata
func SendSuccess(c *gin.Context, code int, data any) {
	c.JSON(code, Response{
		Status: "success",
		Data:   data,
	})
}

// SendSuccessWithMeta sends a successful response with metadata (e.g., pagination)
func SendSuccessWithMeta(c *gin.Context, code int, data any, meta any) {
	c.JSON(code, Response{
		Status: "success",
		Data:   data,
		Meta:   meta,
	})
}

// SendError sends an error response
// errors can be: nil, string, []string, or map[string]string for validation errors
func SendError(c *gin.Context, code int, message string, errors any) {
	c.JSON(code, Response{
		Status:  "error",
		Message: message,
		Errors:  errors,
	})
}
