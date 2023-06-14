package singleton

import (
	"fmt"
	"sync"
)

type Instance struct{}

var (
	once     sync.Once
	instance *Instance
)

func NewInstance() *Instance {
	once.Do(func() {
		instance = &Instance{}
		fmt.Println("Inside")
	})
	fmt.Println("Outside")
	return instance
}

func Start() {
	for i := 0; i < 10; i++ {
		_ = NewInstance()
	}
}
