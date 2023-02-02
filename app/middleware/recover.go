package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				//alarm.Panic(fmt.Sprintf("%s", r))
				c.JSON(http.StatusOK, &Result{
					Code:  10004,
					Data:  nil,
					Error: "服务业务异常请联系技术伙伴",
				})
				return
			}
		}()
		c.Next()
	}
}
