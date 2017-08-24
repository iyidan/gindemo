package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MaxBodyLimit limit the request body
func MaxBodyLimit(b int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPatch ||
			c.Request.Method == http.MethodPost ||
			c.Request.Method == http.MethodPut ||
			c.Request.Method == http.MethodDelete {
			c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, b)
		}
		c.Next()
	}
}
