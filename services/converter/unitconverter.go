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

	WeightUnitMilligram = "Milligram"
	WeightUnitGram      = "Gram"
	WeightUnitKilogram  = "Kilogram"
	WeightUnitOunce     = "Ounce"
	WeightUnitPound     = "Pound"
)

var conversionFactors = map[string]float64{
	// Length
	// Millimeter
	"MillimeterToCentimeter": 0.1,
	"MillimeterToMeter":      0.001,
	"MillimeterToKilometer":  0.000001,
	"MillimeterToInch":       1 / 25.4,
	"MillimeterToFoot":       1 / 304.8,
	"MillimeterToYard":       1 / 914.4,
	"MillimeterToMile":       1.0 / 1609344,

	// Centimeter
	"CentimeterToMillimeter": 10.0,
	"CentimeterToMeter":      0.01,
	"CentimeterToKilometer":  0.00001,
	"CentimeterToInch":       1 / 2.54,
	"CentimeterToFoot":       1 / 30.48,
	"CentimeterToYard":       1 / 91.44,
	"CentimeterToMile":       1 / 160934.4,

	// Meter
	"MeterToMillimeter": 1000,
	"MeterToCentimeter": 100,
	"MeterToKilometer":  0.001,
	"MeterToInch":       39.3701,
	"MeterToFoot":       3.28084,
	"MeterToYard":       1.09361,
	"MeterToMile":       0.000621371,

	// Kilometer
	"KilometerToMillimeter": 1000000,
	"KilometerToCentimeter": 100000,
	"KilometerToMeter":      1000,
	"KilometerToInch":       39370.1,
	"KilometerToFoot":       3280.84,
	"KilometerToYard":       1093.61,
	"KilometerToMile":       0.621371,

	// Inch
	"InchToMillimeter": 25.4,
	"InchToCentimeter": 2.54,
	"InchToMeter":      0.0254,
	"InchToKilometer":  0.0000254,
	"InchToFoot":       1.0 / 12,
	"InchToYard":       1.0 / 36,
	"InchToMile":       1.0 / 63360,

	// Foot
	"FootToMillimeter": 304.8,
	"FootToCentimeter": 30.48,
	"FootToMeter":      0.3048,
	"FootToKilometer":  0.0003048,
	"FootToInch":       12,
	"FootToYard":       1.0 / 3,
	"FootToMile":       1.0 / 5280,

	// Yard
	"YardToMillimeter": 914.4,
	"YardToCentimeter": 91.44,
	"YardToMeter":      0.9144,
	"YardToKilometer":  0.0009144,
	"YardToInch":       36,
	"YardToFoot":       3,
	"YardToMile":       1.0 / 1760,

	// Mile
	"MileToMillimeter": 1609344,
	"MileToCentimeter": 160934.4,
	"MileToMeter":      1609.344,
	"MileToKilometer":  1.609344,
	"MileToInch":       63360,
	"MileToFoot":       5280,
	"MileToYard":       1760,

	// Weight
	// Milligram
	"MilligramToGram":     0.001,
	"MilligramToKilogram": 0.000001,
	"MilligramToOunce":    1.0 / 28349.5,
	"MilligramToPound":    1.0 / 453592,

	// Gram
	"GramToMilligram": 1000,
	"GramToKilogram":  0.001,
	"GramToOunce":     1.0 / 28.3495,
	"GramToPound":     1.0 / 453.592,

	// Kilogram
	"KilogramToMilligram": 1000000,
	"KilogramToGram":      1000,
	"KilogramToOunce":     35.274,
	"KilogramToPound":     2.20462,

	// Ounce
	"OunceToMilligram": 28349.5,
	"OunceToGram":      28.3495,
	"OunceToKilogram":  0.0283495,
	"OunceToPound":     1.0 / 16,

	// Pound
	"PoundToMilligram": 453592,
	"PoundToGram":      453.592,
	"PoundToKilogram":  0.453592,
	"PoundToOunce":     16,
}

type UnitConverter struct {
	strategyMap map[string]ConvertStrategy
}

func NewUnitConverter() *UnitConverter {
	strategyMap := make(map[string]ConvertStrategy)

	for key, factor := range conversionFactors {
		strategyMap[key] = func(value float64) float64 {
			return value * factor
		}
	}

	return &UnitConverter{
		strategyMap: strategyMap,
	}
}

func (s *UnitConverter) ListLengthUnits() []string {
	return []string{
		LengthUnitMillimeter,
		LengthUnitCentimeter,
		LengthUnitMeter,
		LengthUnitKilometer,
		LengthUnitInch,
		LengthUnitFoot,
		LengthUnitYard,
		LengthUnitMile,
	}
}

func (s *UnitConverter) ListWeightUnits() []string {
	return []string{
		WeightUnitMilligram,
		WeightUnitGram,
		WeightUnitKilogram,
		WeightUnitOunce,
		WeightUnitPound,
	}
}

func (s *UnitConverter) Convert(value float64, unitFrom string, unitTo string) float64 {
	k := generateStrategyMapKey(unitFrom, unitTo)

	strategy, ok := s.strategyMap[k]
	if !ok {
		return 0
	}

	return strategy(value)
}
