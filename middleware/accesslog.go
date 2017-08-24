package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/iyidan/gindemo/conf"
)

// Accesslog like nginx access log format
// log_format
// '$remote_addr [$time_iso8601] $status $body_bytes_sent '
// '"$request_method $scheme://$server_name$request_uri" '
// '"$http_referer" "$http_user_agent" "$http_x_forwarded_for" '
// '"$request_time" "$upstream_addr" "$upstream_status" "$upstream_response_time" "$proxy_host"';
func Accesslog() gin.HandlerFunc {

	// @todo log into file
	w := os.Stderr
	serverAddr := conf.String("addr")

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		url := c.Request.URL.String()

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()

		fmt.Fprintf(w, "%s [%s] %d %d \"%s %s\" \"%s\" \"%s\" \"%s\" \"%s\" \"%s\"\n",
			c.ClientIP(),
			end.Format("2006-01-02T15:04:05Z07:00"),
			c.Writer.Status(),
			c.Writer.Size(),
			c.Request.Method,
			url,
			c.Request.Referer(),
			c.Request.UserAgent(),
			c.Request.Header.Get("X-Forwarded-For"),
			end.Sub(start),
			serverAddr)
	}
}
