package services

import "github.com/prometheus/client_golang/prometheus"

var (
	URLShortedBySVC = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "url_shorted_by_service",
			Help: "total number of URL shorted by service",
		},
		[]string{"url"},
	)

	URLShortedByCache = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "url_shorted_by_cache",
			Help: "total number of URL shorted by cache",
		},
		[]string{"url"},
	)
)

func init() {
	prometheus.Register(URLShortedBySVC)
	prometheus.Register(URLShortedByCache)
}
