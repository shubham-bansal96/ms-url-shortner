package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func NewBaseContoller() *BaseController {
	return &BaseController{}
}

func (bc *BaseController) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ping successful")
}
