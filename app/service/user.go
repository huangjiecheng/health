package service

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

var (
	IsArrive    = true
	UserDemoVal UserDemo
)

type UserDemo struct {
	Num float64 `json:"num"`
}

func PrintVal() {
	v := &UserDemo{}
	str := "{\n      \"num\": 123    }"
	_ = json.Unmarshal([]byte(str), v)
	fmt.Printf("%+v,,,,,,%+v", v, reflect.TypeOf(v))
	fmt.Println(time.Now().Format("2006-01-02_15:04:05"))
	//fmt.Printf("打印user：%v, Num: %v\n", IsArrive, UserDemoVal)
}
