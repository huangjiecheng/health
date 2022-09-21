package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct{}

func (c BaseController) Success(ctx *gin.Context) {
	ctx.String(http.StatusOK, "success")
}
