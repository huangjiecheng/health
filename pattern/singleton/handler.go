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
		fmt.Println(fmt.Sprintf("index: %d 懒加载首次初始化实例", i))

	})
	fmt.Println(fmt.Sprintf("index: %d 获取到唯一实例", i))
	return instance
}

func Start() {
	fmt.Println("======单例模式启动======")
	for i := 0; i < 20; i++ {
		go func(x int) {
			_ = NewInstance(x)
		}(i)
	}
	time.Sleep(2 * time.Second)
}
