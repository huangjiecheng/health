package sedan

import "fmt"

// Tire 是SUV汽车的轮胎类型
type Tire struct{}

// Roll 实现了SUV轮胎的滚动方法
func (t *Tire) Roll() {
	fmt.Println("轿车轮胎滚动")
}
