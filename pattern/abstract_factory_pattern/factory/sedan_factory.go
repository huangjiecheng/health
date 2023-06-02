package factory

import (
	"fmt"
	"pattern/abstract_factory_pattern/car"
	"pattern/abstract_factory_pattern/car/sedan"
)

// SedanFactory 是轿车工厂类型
type SedanFactory struct{}

// NewSedanFactory 创建一个新的轿车工厂
func NewSedanFactory() Factory {
	fmt.Println("【轿车】工厂建立。。。")
	return &SedanFactory{}
}

// CreateEngine 创建轿车引擎
func (f *SedanFactory) CreateEngine() car.Engine {
	fmt.Println("【轿车】工厂生产引擎。。。")
	return &sedan.Engine{}
}

// CreateTire 创建轿车轮胎
func (f *SedanFactory) CreateTire() car.Tire {
	fmt.Println("【轿车】工厂生产轮胎。。。")
	return &sedan.Tire{}
}
