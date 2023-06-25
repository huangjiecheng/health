package singleton

import (
	"fmt"
	"sync"
	"time"
)

type Instance struct{}

var (
	once     sync.Once
	instance *Instance
)

func NewInstance(i int) *Instance {
	once.Do(func() {
		instance = &Instance{}
		fmt.Println(fmt.Sprintf("index: %d 懒加载初始化", i))

	})
	fmt.Println(fmt.Sprintf("index: %d 获取到旧的实例", i))
	return instance
}

func Start() {
	fmt.Println("======单例模式启动======")
	for i := 0; i < 10; i++ {
		go func() {
			_ = NewInstance(i)
		}()
	}
	time.Sleep(10 * time.Second)
}
