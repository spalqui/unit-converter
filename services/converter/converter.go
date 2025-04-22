package converter

type ConvertStrategy func(value float64) float64

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

	TemperatureUnitCelsius    = "Celsius"
	TemperatureUnitFahrenheit = "Fahrenheit"
	TemperatureUnitKelvin     = "Kelvin"
)

var conversionFactors = map[string]float64{
	// Length
	// Millimeter
	generateStrategyMapKey(LengthUnitMillimeter, LengthUnitMillimeter): 1.0,
	generateStrategyMapKey(LengthUnitMillimeter, LengthUnitCentimeter): 0.1,
	generateStrategyMapKey(LengthUnitMillimeter, LengthUnitMeter):      0.001,
	generateStrategyMapKey(LengthUnitMillimeter, LengthUnitKilometer):  0.000001,
	generateStrategyMapKey(LengthUnitMillimeter, LengthUnitInch):       1 / 25.4,
	generateStrategyMapKey(LengthUnitMillimeter, LengthUnitFoot):       1 / 304.8,
	generateStrategyMapKey(LengthUnitMillimeter, LengthUnitYard):       1 / 914.4,
	generateStrategyMapKey(LengthUnitMillimeter, LengthUnitMile):       1.0 / 1609344,

	// Centimeter
	generateStrategyMapKey(LengthUnitCentimeter, LengthUnitCentimeter): 1.0,
	generateStrategyMapKey(LengthUnitCentimeter, LengthUnitMillimeter): 10.0,
	generateStrategyMapKey(LengthUnitCentimeter, LengthUnitMeter):      0.01,
	generateStrategyMapKey(LengthUnitCentimeter, LengthUnitKilometer):  0.00001,
	generateStrategyMapKey(LengthUnitCentimeter, LengthUnitInch):       1 / 2.54,
	generateStrategyMapKey(LengthUnitCentimeter, LengthUnitFoot):       1 / 30.48,
	generateStrategyMapKey(LengthUnitCentimeter, LengthUnitYard):       1 / 91.44,
	generateStrategyMapKey(LengthUnitCentimeter, LengthUnitMile):       1 / 160934.4,

	// Meter
	generateStrategyMapKey(LengthUnitMeter, LengthUnitMeter):      1.0,
	generateStrategyMapKey(LengthUnitMeter, LengthUnitMillimeter): 1000,
	generateStrategyMapKey(LengthUnitMeter, LengthUnitCentimeter): 100,
	generateStrategyMapKey(LengthUnitMeter, LengthUnitKilometer):  0.001,
	generateStrategyMapKey(LengthUnitMeter, LengthUnitInch):       39.3701,
	generateStrategyMapKey(LengthUnitMeter, LengthUnitFoot):       3.28084,
	generateStrategyMapKey(LengthUnitMeter, LengthUnitYard):       1.09361,
	generateStrategyMapKey(LengthUnitMeter, LengthUnitMile):       0.000621371,

	// Kilometer
	generateStrategyMapKey(LengthUnitKilometer, LengthUnitKilometer):  1.0,
	generateStrategyMapKey(LengthUnitKilometer, LengthUnitMillimeter): 1000000,
	generateStrategyMapKey(LengthUnitKilometer, LengthUnitCentimeter): 100000,
	generateStrategyMapKey(LengthUnitKilometer, LengthUnitMeter):      1000,
	generateStrategyMapKey(LengthUnitKilometer, LengthUnitInch):       39370.1,
	generateStrategyMapKey(LengthUnitKilometer, LengthUnitFoot):       3280.84,
	generateStrategyMapKey(LengthUnitKilometer, LengthUnitYard):       1093.61,
	generateStrategyMapKey(LengthUnitKilometer, LengthUnitMile):       0.621371,

	// Inch
	generateStrategyMapKey(LengthUnitInch, LengthUnitInch):       1.0,
	generateStrategyMapKey(LengthUnitInch, LengthUnitMillimeter): 25.4,
	generateStrategyMapKey(LengthUnitInch, LengthUnitCentimeter): 2.54,
	generateStrategyMapKey(LengthUnitInch, LengthUnitMeter):      0.0254,
	generateStrategyMapKey(LengthUnitInch, LengthUnitKilometer):  0.0000254,
	generateStrategyMapKey(LengthUnitInch, LengthUnitFoot):       1.0 / 12,
	generateStrategyMapKey(LengthUnitInch, LengthUnitYard):       1.0 / 36,
	generateStrategyMapKey(LengthUnitInch, LengthUnitMile):       1.0 / 63360,

	// Foot
	generateStrategyMapKey(LengthUnitFoot, LengthUnitFoot):       1.0,
	generateStrategyMapKey(LengthUnitFoot, LengthUnitMillimeter): 304.8,
	generateStrategyMapKey(LengthUnitFoot, LengthUnitCentimeter): 30.48,
	generateStrategyMapKey(LengthUnitFoot, LengthUnitMeter):      0.3048,
	generateStrategyMapKey(LengthUnitFoot, LengthUnitKilometer):  0.0003048,
	generateStrategyMapKey(LengthUnitFoot, LengthUnitInch):       12,
	generateStrategyMapKey(LengthUnitFoot, LengthUnitYard):       1.0 / 3,
	generateStrategyMapKey(LengthUnitFoot, LengthUnitMile):       1.0 / 5280,

	// Yard
	generateStrategyMapKey(LengthUnitYard, LengthUnitYard):       1.0,
	generateStrategyMapKey(LengthUnitYard, LengthUnitMillimeter): 914.4,
	generateStrategyMapKey(LengthUnitYard, LengthUnitCentimeter): 91.44,
	generateStrategyMapKey(LengthUnitYard, LengthUnitMeter):      0.9144,
	generateStrategyMapKey(LengthUnitYard, LengthUnitKilometer):  0.0009144,
	generateStrategyMapKey(LengthUnitYard, LengthUnitInch):       36,
	generateStrategyMapKey(LengthUnitYard, LengthUnitFoot):       3,
	generateStrategyMapKey(LengthUnitYard, LengthUnitMile):       1.0 / 1760,

	// Mile
	generateStrategyMapKey(LengthUnitMile, LengthUnitMile):       1.0,
	generateStrategyMapKey(LengthUnitMile, LengthUnitMillimeter): 1609344,
	generateStrategyMapKey(LengthUnitMile, LengthUnitCentimeter): 160934.4,
	generateStrategyMapKey(LengthUnitMile, LengthUnitMeter):      1609.344,
	generateStrategyMapKey(LengthUnitMile, LengthUnitKilometer):  1.609344,
	generateStrategyMapKey(LengthUnitMile, LengthUnitInch):       63360,
	generateStrategyMapKey(LengthUnitMile, LengthUnitFoot):       5280,
	generateStrategyMapKey(LengthUnitMile, LengthUnitYard):       1760,

	// Weight
	// Milligram
	generateStrategyMapKey(WeightUnitMilligram, WeightUnitMilligram): 1.0,
	generateStrategyMapKey(WeightUnitMilligram, WeightUnitGram):      0.001,
	generateStrategyMapKey(WeightUnitMilligram, WeightUnitKilogram):  0.000001,
	generateStrategyMapKey(WeightUnitMilligram, WeightUnitOunce):     1.0 / 28349.5,
	generateStrategyMapKey(WeightUnitMilligram, WeightUnitPound):     1.0 / 453592,

	// Gram
	generateStrategyMapKey(WeightUnitGram, WeightUnitGram):      1.0,
	generateStrategyMapKey(WeightUnitGram, WeightUnitMilligram): 1000,
	generateStrategyMapKey(WeightUnitGram, WeightUnitKilogram):  0.001,
	generateStrategyMapKey(WeightUnitGram, WeightUnitOunce):     1.0 / 28.3495,
	generateStrategyMapKey(WeightUnitGram, WeightUnitPound):     1.0 / 453.592,

	// Kilogram
	generateStrategyMapKey(WeightUnitKilogram, WeightUnitKilogram):  1.0,
	generateStrategyMapKey(WeightUnitKilogram, WeightUnitMilligram): 1000000,
	generateStrategyMapKey(WeightUnitKilogram, WeightUnitGram):      1000,
	generateStrategyMapKey(WeightUnitKilogram, WeightUnitOunce):     35.274,
	generateStrategyMapKey(WeightUnitKilogram, WeightUnitPound):     2.20462,

	// Ounce
	generateStrategyMapKey(WeightUnitOunce, WeightUnitOunce):     1.0,
	generateStrategyMapKey(WeightUnitOunce, WeightUnitMilligram): 28349.5,
	generateStrategyMapKey(WeightUnitOunce, WeightUnitGram):      28.3495,
	generateStrategyMapKey(WeightUnitOunce, WeightUnitKilogram):  0.0283495,
	generateStrategyMapKey(WeightUnitOunce, WeightUnitPound):     1.0 / 16,

	// Pound
	generateStrategyMapKey(WeightUnitPound, WeightUnitPound):     1.0,
	generateStrategyMapKey(WeightUnitPound, WeightUnitMilligram): 453592,
	generateStrategyMapKey(WeightUnitPound, WeightUnitGram):      453.592,
	generateStrategyMapKey(WeightUnitPound, WeightUnitKilogram):  0.453592,
	generateStrategyMapKey(WeightUnitPound, WeightUnitOunce):     16,
}

var temperatureConversionStrategies = map[string]ConvertStrategy{
	// Temperature
	generateStrategyMapKey(TemperatureUnitCelsius, TemperatureUnitCelsius): func(value float64) float64 {
		return value
	},
	generateStrategyMapKey(TemperatureUnitCelsius, TemperatureUnitFahrenheit): func(value float64) float64 {
		return (value * 9 / 5) + 32
	},
	generateStrategyMapKey(TemperatureUnitCelsius, TemperatureUnitKelvin): func(value float64) float64 {
		return value + 273.15
	},
	generateStrategyMapKey(TemperatureUnitFahrenheit, TemperatureUnitFahrenheit): func(value float64) float64 {
		return value
	},
	generateStrategyMapKey(TemperatureUnitFahrenheit, TemperatureUnitCelsius): func(value float64) float64 {
		return (value - 32) * 5 / 9
	},
	generateStrategyMapKey(TemperatureUnitFahrenheit, TemperatureUnitKelvin): func(value float64) float64 {
		return (value-32)*5/9 + 273.15
	},
	generateStrategyMapKey(TemperatureUnitKelvin, TemperatureUnitKelvin): func(value float64) float64 {
		return value
	},
	generateStrategyMapKey(TemperatureUnitKelvin, TemperatureUnitCelsius): func(value float64) float64 {
		return value - 273.15
	},
	generateStrategyMapKey(TemperatureUnitKelvin, TemperatureUnitFahrenheit): func(value float64) float64 {
		return (value-273.15)*9/5 + 32
	},
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

	for key, strategy := range temperatureConversionStrategies {
		strategyMap[key] = strategy
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

func (s *UnitConverter) ListTemperatureUnits() []string {
	return []string{
		TemperatureUnitCelsius,
		TemperatureUnitFahrenheit,
		TemperatureUnitKelvin,
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

func generateStrategyMapKey(unitFrom, unitTo string) string {
	return unitFrom + "To" + unitTo
}
