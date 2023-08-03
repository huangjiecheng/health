package main

import (
	"pattern/abstract_factory_pattern"
	"pattern/singleton"
	"pattern/strategy_pattern"
)

func main() {
	singleton.Start()

	abstract_factory_pattern.Start()

	strategy_pattern.Start()
}
