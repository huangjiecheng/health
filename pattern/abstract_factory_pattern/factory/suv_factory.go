package factory

import (
	"fmt"
	"pattern/abstract_factory_pattern/car"
	"pattern/abstract_factory_pattern/car/suv"
)

// SUVFactory 是SUV工厂类型
type SUVFactory struct{}

// NewSUVFactory 创建一个新的SUV工厂
func NewSUVFactory() Factory {
	fmt.Println("【SUV】工厂建立")
	return &SUVFactory{}
}

// CreateEngine 创建SUV引擎
func (f *SUVFactory) CreateEngine() car.Engine {
	fmt.Println("【SUV】工厂生产引擎。。。")
	return &suv.Engine{}
}

// CreateTire 创建SUV轮胎
func (f *SUVFactory) CreateTire() car.Tire {
	fmt.Println("【SUV】工厂生产轮胎。。。")
	return &suv.Tire{}
}
