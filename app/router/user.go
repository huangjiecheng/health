package router

import (
	"github.com/gin-gonic/gin"
	"health/controller"
	"health/middleware"
)

type userRouter struct{}

func (u *userRouter) registerApi(r *gin.Engine) {
	v1 := r.Group("/v1/users")
	v1.GET("")
	v1.GET("/:id", middleware.Auth(), controller.UserController{}.Get)
	v1.GET("/api/project/:project_id/users", controller.UserController{}.List)
}
