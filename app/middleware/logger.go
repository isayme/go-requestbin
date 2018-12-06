package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	logger "github.com/isayme/go-logger"
	"github.com/isayme/go-requestbin/app/constant"
)

// Logger log middlware
func Logger(c *gin.Context) {
	// Start timer
	start := time.Now()

	// Process request
	c.Next()

	// Stop timer
	end := time.Now()
	latency := end.Sub(start)

	method := c.Request.Method
	url := c.Request.RequestURI
	statusCode := c.Writer.Status()
	userAgent := c.GetHeader(constant.HeaderUserAgent)
	ip := c.ClientIP()

	logger.Infow("", "method", method, "url", url, "status", statusCode, "userAgent", userAgent, "ip", ip, "latency", latency.String())
}
