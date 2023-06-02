package strategy_pattern

import (
	"fmt"
	"pattern/strategy_pattern/order"
	"pattern/strategy_pattern/strategy"
)

func Start() {
	fmt.Println("======策略模式启动======")
	var (
		ctx      = &order.Context{}
		amount   = 120.0
		quantity = 5
	)

	fmt.Printf("【单价】%.2f 【数量】：%d 【总额】：%.2f \n", amount, quantity, amount*float64(quantity))
	// 设置满减策略
	ctx.SetStrategy(&strategy.FullReductionStrategy{
		Threshold: 100,
		Reduction: 20,
	})

	fmt.Printf("【每满100-20】：%.2f\n", ctx.CalculateDiscount(amount, quantity))

	// 设置折扣策略
	ctx.SetStrategy(&strategy.DiscountStrategy{
		Percentage: 0.1,
	})

	fmt.Printf("【原价打9折】：%.2f\n", ctx.CalculateDiscount(amount, quantity))

	// 设置第二件半价策略
	ctx.SetStrategy(&strategy.SecondHalfPriceStrategy{})

	fmt.Printf("【第二件半价】：%.2f\n", ctx.CalculateDiscount(amount, quantity))
}
