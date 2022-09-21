package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 定义ip白名单
		whiteList := []string{
			"127.0.0.1",
		}

		ip := ctx.ClientIP()

		flag := false

		for _, host := range whiteList {
			if ip == host {
				flag = true
				break
			}
		}

		if !flag {
			ctx.AbortWithStatusJSON(http.StatusNetworkAuthenticationRequired, fmt.Sprintf("your ip is not trusted: %s", ip))
		}

	}
}
