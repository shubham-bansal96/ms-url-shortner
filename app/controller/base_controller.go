package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ms-url-shortner/app/logging"
	"github.com/ms-url-shortner/app/model"
	"github.com/ms-url-shortner/app/services"
	"github.com/ms-url-shortner/app/utils"
)

type BaseController struct {
	ShortenURLService services.IShortenUrl
}

// NewBaseContoller creates a new BaseController with the given URL shortening service
func NewBaseContoller(sus services.IShortenUrl) *BaseController {
	return &BaseController{
		ShortenURLService: sus,
	}
}

// Ping responds with a simple success message to verify the service is running
func (bc *BaseController) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ping successful")
}

// HandleURLShortner binds the JSON request, validates the URL, and returns the shortened URL
func (bc *BaseController) HandleURLShortner(ctx *gin.Context) {
	lw := logging.LogForFunc()

	var requestObj *model.URLDTO
	if err := ctx.ShouldBindJSON(&requestObj); err != nil {
		lw.WithField("error", "error while binding JSON request").Error(err.Error())
		utils.RendorJson(ctx, nil, http.StatusBadRequest, model.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := requestObj.Validate(); err != nil {
		utils.RendorJson(ctx, nil, *err.ErrorCode, err)
		return
	}

	data := bc.ShortenURLService.ShortURL(ctx, *requestObj.URL)

	utils.RendorJson(ctx, data, http.StatusOK, nil)
	return
}
