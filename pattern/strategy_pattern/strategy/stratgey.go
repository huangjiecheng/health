package strategy

type Strategy interface {
	CalculateDiscount(amount float64, quantity int) float64
}
