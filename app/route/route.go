package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ms-url-shortner/app/config"
	"github.com/ms-url-shortner/app/controller"
	"github.com/ms-url-shortner/app/services"
)

func Initialize(router *gin.Engine) {
	ctrl := controller.NewBaseContoller(services.NewShortenURLService(services.NewUidService()))

	appGroup := router.Group("/" + config.Config.MSName)
	appGroup.GET(testEndPoint, ctrl.Ping)
	appGroup.POST(getShortURLEndPoint, ctrl.HandleURLRequest)
}
