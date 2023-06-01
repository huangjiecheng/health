package factory

import (
	car2 "pattern/abstract_factory/car"
	suv2 "pattern/abstract_factory/car/suv"
)

// SUVFactory 是SUV工厂类型
type SUVFactory struct{}

// NewSUVFactory 创建一个新的SUV工厂
func NewSUVFactory() Factory {
	return &SUVFactory{}
}

// CreateEngine 创建SUV引擎
func (f *SUVFactory) CreateEngine() car2.Engine {
	return &suv2.Engine{}
}

// CreateTire 创建SUV轮胎
func (f *SUVFactory) CreateTire() car2.Tire {
	return &suv2.Tire{}
}
