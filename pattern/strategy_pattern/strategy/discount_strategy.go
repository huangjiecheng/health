package strategy

type DiscountStrategy struct {
	Percentage float64
}

func (s *DiscountStrategy) CalculateDiscount(amount float64, quantity int) float64 {
	total := amount * float64(quantity)
	return total * (1 - s.Percentage)
}
