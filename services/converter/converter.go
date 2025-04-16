package converter

type Converter interface {
	Convert(value float64, unitFrom string, unitTo string)
}

type ConvertStrategy func(value float64) float64

func generateStrategyMapKey(unitFrom, unitTo string) string {
	return unitFrom + "To" + unitTo
}
