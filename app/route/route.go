package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ms-url-shortner/app/config"
	"github.com/ms-url-shortner/app/controller"
)

func Initialize(router *gin.Engine) {
	ctrl := controller.NewBaseContoller()

	appGroup := router.Group("/" + config.Config.MSName)
	appGroup.GET(testEndPoint, ctrl.Ping)
}
