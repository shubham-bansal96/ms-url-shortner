package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{"path"},
)

func init() {
	prometheus.Register(totalRequests)
}

func MetricMiddleware(handler func(c *gin.Context)) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		totalRequests.WithLabelValues("ping").Inc()
		handler(ctx)
	}
}
