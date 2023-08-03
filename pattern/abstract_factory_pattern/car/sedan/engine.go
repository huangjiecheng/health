package sedan

import "fmt"

// Engine 是轿车的引擎类型
type Engine struct{}

// Start 实现了轿车引擎的启动方法
func (e *Engine) Start() {
	fmt.Println("【轿车】启动引擎")
}

// Stop 实现了轿车引擎的停止方法
func (e *Engine) Stop() {
	fmt.Println("【轿车】停止引擎")
}
