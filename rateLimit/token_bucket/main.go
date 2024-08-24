package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

var r *gin.Engine

func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {
	r = gin.Default()
	r.GET("/", RateLimitMiddleware(time.Second*3, 5), func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.Run(":8080")
}
