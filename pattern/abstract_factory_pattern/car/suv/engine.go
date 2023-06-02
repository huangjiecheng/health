package suv

import "fmt"

// Engine 是SUV汽车的引擎类型
type Engine struct{}

// Start 实现了SUV引擎的启动方法
func (e *Engine) Start() {
	fmt.Println("启动SUV引擎")
}

// Stop 实现了SUV引擎的停止方法
func (e *Engine) Stop() {
	fmt.Println("停止SUV引擎")
}
