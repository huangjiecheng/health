package order

import "pattern/strategy_pattern/strategy"

type Context struct {
	strategy strategy.Strategy
}

func (c *Context) SetStrategy(strategy strategy.Strategy) {
	c.strategy = strategy
}

func (c *Context) CalculateDiscount(amount float64, quantity int) float64 {
	if c.strategy != nil {
		return c.strategy.CalculateDiscount(amount, quantity)
	}
	return amount
}
