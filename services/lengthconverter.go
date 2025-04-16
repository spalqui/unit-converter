package services

type LengthConverter struct {
	strategy ConvertStrategy
}

func NewLengthConverter() *LengthConverter {
	return &LengthConverter{}
}

func (s *LengthConverter) SetStrategy(strategy ConvertStrategy) {
	s.strategy = strategy
}

func (s *LengthConverter) Convert(value float64) float64 {
	return s.strategy(value)
}

func MillimeterToCentimeter(value float64) float64 {
	return value / 10
}
