package middleware

import (
	"BlogService/pkg/app"
	"BlogService/pkg/errcode"
	"BlogService/pkg/limiter"

	"github.com/gin-gonic/gin"
)

func Limiter(l limiter.InterfaceLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
