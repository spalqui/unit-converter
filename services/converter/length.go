package converter

const (
	LengthUnitMillimeter = "Millimeter"
	LengthUnitCentimeter = "Centimeter"
	LengthUnitMeter      = "Meter"
	LengthUnitKilometer  = "Kilometer"
	LengthUnitInch       = "Inch"
	LengthUnitFoot       = "Foot"
	LengthUnitYard       = "Yard"
	LengthUnitMile       = "Mile"
)

type LengthConverter struct {
	strategyMap map[string]ConvertStrategy
}

func NewLengthConverter() *LengthConverter {
	strategyMap := make(map[string]ConvertStrategy)

	strategyMap[generateStrategyMapKey(LengthUnitMillimeter, LengthUnitCentimeter)] = millimeterToCentimeter
	strategyMap[generateStrategyMapKey(LengthUnitMillimeter, LengthUnitMeter)] = millimeterToMeter

	strategyMap[generateStrategyMapKey(LengthUnitCentimeter, LengthUnitMillimeter)] = centimeterToMillimeter

	return &LengthConverter{
		strategyMap: strategyMap,
	}
}

func (s *LengthConverter) Convert(value float64, unitFrom string, unitTo string) float64 {
	k := generateStrategyMapKey(unitFrom, unitTo)

	strategy, ok := s.strategyMap[k]
	if !ok {
		return 0
	}

	return strategy(value)
}

func millimeterToCentimeter(value float64) float64 {
	return value / 10
}

func millimeterToMeter(value float64) float64 {
	return value * 0.001
}

func millimeterToKilometer(value float64) float64 {
	return value * 0.000001
}

func millimeterToInch(value float64) float64 {
	return value / 25.4
}

func centimeterToMillimeter(value float64) float64 {
	return value * 10
}
