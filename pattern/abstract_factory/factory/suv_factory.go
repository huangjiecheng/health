package factory

import (
	"pattern/abstract_factory/car"
	"pattern/abstract_factory/car/suv"
)

// SUVFactory 是SUV工厂类型
type SUVFactory struct{}

// NewSUVFactory 创建一个新的SUV工厂
func NewSUVFactory() Factory {
	return &SUVFactory{}
}

// CreateEngine 创建SUV引擎
func (f *SUVFactory) CreateEngine() car.Engine {
	return &suv.Engine{}
}

// CreateTire 创建SUV轮胎
func (f *SUVFactory) CreateTire() car.Tire {
	return &suv.Tire{}
}
