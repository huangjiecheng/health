package factory

import (
	car2 "pattern/abstract_factory/car"
	sedan2 "pattern/abstract_factory/car/sedan"
)

// SedanFactory 是轿车工厂类型
type SedanFactory struct{}

// NewSedanFactory 创建一个新的轿车工厂
func NewSedanFactory() Factory {
	return &SedanFactory{}
}

// CreateEngine 创建轿车引擎
func (f *SedanFactory) CreateEngine() car2.Engine {
	return &sedan2.Engine{}
}

// CreateTire 创建轿车轮胎
func (f *SedanFactory) CreateTire() car2.Tire {
	return &sedan2.Tire{}
}
