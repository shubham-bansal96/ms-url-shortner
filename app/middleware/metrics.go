package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	totalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of get requests.",
		},
		[]string{"path"},
	)

	httpRequestDuration = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "http_requests_duration_seconds",
			Help:       "Time taken to serve the request by path and method.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"path"},
	)

	httpRequestInFlight = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_requests_in_flight",
			Help: "Number of requests currently being served by path and method.",
		},
		[]string{"path"},
	)
)

func init() {
	prometheus.Register(totalRequests)
	prometheus.Register(httpRequestDuration)
	prometheus.Register(httpRequestInFlight)
}

// func MetricMiddleware(handler func(c *gin.Context)) func(c *gin.Context) {
// 	return func(ctx *gin.Context) {
// 		totalRequests.WithLabelValues("ping").Inc()
// 		handler(ctx)
// 	}
// }

func MetricMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		totalRequests.WithLabelValues(ctx.Request.URL.Path).Inc()
		httpRequestInFlight.WithLabelValues(ctx.Request.URL.Path).Inc()
		httpRequestDuration := prometheus.NewTimer(httpRequestDuration.WithLabelValues(ctx.Request.URL.Path))

		defer func() {
			httpRequestDuration.ObserveDuration()
			httpRequestInFlight.WithLabelValues(ctx.Request.URL.Path).Dec()
		}()
		ctx.Next()
	})
}
