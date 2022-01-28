package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/ms-url-shortner/app/model"
)

func RendorJson(ctx *gin.Context, data interface{}, statusCode int, err *model.Error) {
	responseDTO := &model.ResponseDTO{
		Data:  data,
		Error: err,
	}

	ctx.JSON(statusCode, responseDTO)
}
