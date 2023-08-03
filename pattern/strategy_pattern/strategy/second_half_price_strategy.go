package strategy

import "math"

type SecondHalfPriceStrategy struct {
}

// CalculateDiscount 第二件半价
func (s *SecondHalfPriceStrategy) CalculateDiscount(amount float64, quantity int) float64 {
	if quantity >= 2 {
		return math.Ceil(float64(quantity)/float64(2)) * amount
	}
	return amount * float64(quantity)
}
