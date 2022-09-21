package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	BaseController
}

// Get 通过ID获取用户信息
func (c UserController) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "User ID:"+ctx.Param("id"))
}

// List 获取用户列表信息
func (c UserController) List(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, "User list")
}
