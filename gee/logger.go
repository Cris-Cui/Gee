package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer	(中间件上半部)
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time  (中间件下半部)
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
