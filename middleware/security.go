package middleware

import "github.com/gin-gonic/gin"

// SecurityHeadersMiddleware adds security-related HTTP headers
func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Prevent MIME type sniffing
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		
		// Prevent clickjacking attacks
		c.Writer.Header().Set("X-Frame-Options", "DENY")
		
		// Enable XSS protection (for older browsers)
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		
		// Enforce HTTPS in production (Strict-Transport-Security)
		// Uncomment in production with HTTPS enabled
		// c.Writer.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		
		// Prevent browser from sending referrer information
		c.Writer.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		
		// Content Security Policy - adjust based on your needs
		c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'")
		
		c.Next()
	}
}
