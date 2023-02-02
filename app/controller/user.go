package controller

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"time"
)

type UserController struct {
	BaseController
}

// Get 通过ID获取用户信息
func (c UserController) Get(ctx *gin.Context) {
	go ForwardSNodesReqStatistical("http://127.0.0.1:8080/api/v1/users/list")
	ctx.JSON(http.StatusOK, "User ID:"+ctx.Param("id"))
}

// GetGoRoutingNum  ss
func (c UserController) GetGoRoutingNum(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, fmt.Sprintf("NumGoroutine【%d】NumCPU【%d】NumCgoCall【%d】",
		runtime.NumGoroutine(), runtime.NumCPU(), runtime.NumCgoCall()))
}

// List 获取用户列表信息
func (c UserController) List(ctx *gin.Context) {
	fmt.Println("hhhahahahahahahahah")
	ctx.JSON(http.StatusBadRequest, "User list")
}

func ForwardSNodesReqStatistical(url string) (respBody []byte, err error) {
	resp, err := forwardReqToMaster(url)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(fmt.Sprintf("ForwardSNodesReqStatistical err: %v", err))
		return nil, err
	}
	// 因为没有错误，即使没有读取body内容，也是需要close的
	return
}

func forwardReqToMaster(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte(`{name:"abc"}`)))
	if err != nil {
		return nil, err
	}
	tsp := &http.Transport{
		MaxIdleConns:          30,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		ResponseHeaderTimeout: time.Duration(40) * time.Second,
	}
	client := &http.Client{Transport: tsp}
	time.Sleep(1 * time.Hour)
	return client.Do(req)
}
