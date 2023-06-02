package strategy

type FullReductionStrategy struct {
	Threshold float64
	Reduction float64
}

func (s *FullReductionStrategy) CalculateDiscount(amount float64, quantity int) float64 {
	total := amount * float64(quantity)
	reduceNum := total / s.Threshold
	if reduceNum > 0 {
		return total - s.Reduction*reduceNum
	}
	return total
}
