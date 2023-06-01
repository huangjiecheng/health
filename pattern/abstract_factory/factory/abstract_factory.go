package factory

import (
	"pattern/abstract_factory/car"
)

// Factory 是抽象工厂接口，定义了创建引擎和轮胎的方法
type Factory interface {
	CreateEngine() car.Engine
	CreateTire() car.Tire
}
