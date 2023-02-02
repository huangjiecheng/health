package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"health/controller"
	"health/middleware"
	"net/http"
)

var (
	user = &userRouter{}
)

func Init() {
	r := gin.Default()
	if gin.Mode() != gin.ReleaseMode {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	r.Use(middleware.LoggerToFile(), middleware.Recover())
	r.Use(gin.Logger())
	v1 := r.Group("/api/v1")
	user.registerApi(r)
	{
		// 注册路由
		v1.GET("/users/:id", middleware.Auth(), controller.UserController{}.Get)
		v1.GET("/users1", middleware.Auth(), controller.UserController{}.GetGoRoutingNum)
		v1.GET("/users/list", middleware.Auth(), controller.UserController{}.List)
		v1.GET("/api/project/:project_id/users", controller.UserController{}.List)
	}

	// v2路由组
	v2 := r.Group("/api/v2")
	{
		// 注册路由
		v2.GET("/user/:id", controller.UserController{}.Get)
		v2.GET("/hjc/:name/*action", func(c *gin.Context) {
			name := c.Param("name")
			action := c.Param("action")
			firstname := c.DefaultQuery("name", "kim") // 获取query中的name，没有的话就为kim
			lastname := c.Query("age")
			message := name + " is: " + action + ",firstname: " + firstname + ",lastname: " + lastname
			c.String(http.StatusOK, message)
		})
	}

	r.Run()
}
