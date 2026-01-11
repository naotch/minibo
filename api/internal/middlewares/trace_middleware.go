package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func AccessTraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		status := c.Writer.Status()
		clientIP := c.ClientIP()

		log.Printf("[TRACE] %3d | %13v | %15s | %-7s %s",
			status, latency, clientIP, method, path,
		)
	}
}
