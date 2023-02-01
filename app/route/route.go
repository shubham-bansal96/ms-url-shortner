package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ms-url-shortner/app/config"
	"github.com/ms-url-shortner/app/controller"
	"github.com/ms-url-shortner/app/middleware"
	"github.com/ms-url-shortner/app/services"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Initialize(router *gin.Engine) {
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	ctrl := controller.NewBaseContoller(services.NewShortenURLService(services.NewUidService()))
	appGroup := router.Group("/" + config.Config.MSName)
	appGroup.GET(testEndPoint, middleware.MetricMiddleware(ctrl.Ping))
	appGroup.POST(getShortURLEndPoint, ctrl.HandleURLShortner)
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
