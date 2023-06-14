package abstract_factory_pattern

import (
	"fmt"
	"pattern/abstract_factory_pattern/factory"
)

func Start() {
	fmt.Println("======抽象工厂模式启动======")
	// 创建SUV工厂
	suvFactory := factory.NewSUVFactory()

	// 创建SUV引擎
	suvEngine := suvFactory.CreateEngine()
	suvEngine.Start()

	// 创建SUV轮胎
	suvTire := suvFactory.CreateTire()
	suvTire.Roll()

	// 创建轿车工厂
	sedanFactory := factory.NewSedanFactory()

	// 创建轿车引擎
	sedanEngine := sedanFactory.CreateEngine()
	sedanEngine.Start()

	// 创建轿车轮胎
	sedanTire := sedanFactory.CreateTire()
	sedanTire.Roll()

	fmt.Println("汽车制造完成。")
}
