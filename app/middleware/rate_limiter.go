package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiter() gin.HandlerFunc {
	limit := rate.NewLimiter(rate.Every(time.Second), 2)
	return gin.HandlerFunc(func(ctx *gin.Context) {
		if !limit.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "limit increased"})
			return
		}
		ctx.Next()
	})
}
