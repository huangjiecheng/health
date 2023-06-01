package abstract_factory

import (
	"fmt"
	factory2 "pattern/abstract_factory/factory"
)

func Start() {
	// 创建SUV工厂
	suvFactory := factory2.NewSUVFactory()

	// 创建SUV引擎
	suvEngine := suvFactory.CreateEngine()
	suvEngine.Start()

	// 创建SUV轮胎
	suvTire := suvFactory.CreateTire()
	suvTire.Roll()

	// 创建轿车工厂
	sedanFactory := factory2.NewSedanFactory()

	// 创建轿车引擎
	sedanEngine := sedanFactory.CreateEngine()
	sedanEngine.Start()

	// 创建轿车轮胎
	sedanTire := sedanFactory.CreateTire()
	sedanTire.Roll()

	fmt.Println("汽车制造完成。")
}
