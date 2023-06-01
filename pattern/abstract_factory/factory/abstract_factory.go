package factory

import (
	car2 "pattern/abstract_factory/car"
)

// Factory 是抽象工厂接口，定义了创建引擎和轮胎的方法
type Factory interface {
	CreateEngine() car2.Engine
	CreateTire() car2.Tire
}
