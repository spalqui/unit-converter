package converter

import (
	"fmt"
	"testing"
)

func TestLengthConverter_Convert(t *testing.T) {
	converter := NewLengthConverter()

	tests := []struct {
		name     string
		value    float64
		unitFrom string
		unitTo   string
		want     float64
	}{
		// Millimeter
		{"Convert 10mm to cm", 10, LengthUnitMillimeter, LengthUnitCentimeter, 1},
		{"Convert 1000mm to m", 1000, LengthUnitMillimeter, LengthUnitMeter, 1},
		{"Convert 1,000,000mm to km", 1_000_000, LengthUnitMillimeter, LengthUnitKilometer, 1},
		{"Convert 100mm to in", 100, LengthUnitMillimeter, LengthUnitInch, 3.93701},
		{"Convert 1700mm to ft", 1700, LengthUnitMillimeter, LengthUnitFoot, 5.57743},
		{"Convert 1500mm to yd", 1500, LengthUnitMillimeter, LengthUnitYard, 1.64042},
		{"Convert 804672mm to mi", 804672, LengthUnitMillimeter, LengthUnitMile, 0.5},
		// Centimeter
		{"Convert 1cm to mm", 1, LengthUnitCentimeter, LengthUnitMillimeter, 10},
		{"Convert 100cm to m", 100, LengthUnitCentimeter, LengthUnitMeter, 1},
		{"Convert 100000cm to km", 100000, LengthUnitCentimeter, LengthUnitKilometer, 1},
		{"Convert 2.54cm to in", 2.54, LengthUnitCentimeter, LengthUnitInch, 1},
		{"Convert 30.48cm to ft", 30.48, LengthUnitCentimeter, LengthUnitFoot, 1},
		{"Convert 91.44cm to yd", 91.44, LengthUnitCentimeter, LengthUnitYard, 1},
		{"Convert 160934.4cm to mi", 160934.4, LengthUnitCentimeter, LengthUnitMile, 1},
		// Meter
		{"Convert 1m to mm", 1, LengthUnitMeter, LengthUnitMillimeter, 1000},
		{"Convert 1m to cm", 1, LengthUnitMeter, LengthUnitCentimeter, 100},
		{"Convert 1m to km", 1, LengthUnitMeter, LengthUnitKilometer, 0.001},
		{"Convert 1m to in", 1, LengthUnitMeter, LengthUnitInch, 39.3701},
		{"Convert 1m to ft", 1, LengthUnitMeter, LengthUnitFoot, 3.28084},
		{"Convert 1m to yd", 1, LengthUnitMeter, LengthUnitYard, 1.09361},
		{"Convert 1m to mi", 1, LengthUnitMeter, LengthUnitMile, 0.000621371},
		{"Convert 1000m to km", 1000, LengthUnitMeter, LengthUnitKilometer, 1},
		{"Convert 0.001m to mm", 0.001, LengthUnitMeter, LengthUnitMillimeter, 1},

		// Kilometer
		{"Convert 1km to mm", 1, LengthUnitKilometer, LengthUnitMillimeter, 1000000},
		{"Convert 1km to cm", 1, LengthUnitKilometer, LengthUnitCentimeter, 100000},
		{"Convert 1km to m", 1, LengthUnitKilometer, LengthUnitMeter, 1000},
		{"Convert 1km to in", 1, LengthUnitKilometer, LengthUnitInch, 39370.1},
		{"Convert 1km to ft", 1, LengthUnitKilometer, LengthUnitFoot, 3280.84},
		{"Convert 1km to yd", 1, LengthUnitKilometer, LengthUnitYard, 1093.61},
		{"Convert 1km to mi", 1, LengthUnitKilometer, LengthUnitMile, 0.621371},

		// Inch
		{"Convert 1in to mm", 1, LengthUnitInch, LengthUnitMillimeter, 25.4},
		{"Convert 1in to cm", 1, LengthUnitInch, LengthUnitCentimeter, 2.54},
		{"Convert 1in to m", 1, LengthUnitInch, LengthUnitMeter, 0.0254},
		{"Convert 1in to km", 1, LengthUnitInch, LengthUnitKilometer, 0.0000254},
		{"Convert 1in to ft", 1, LengthUnitInch, LengthUnitFoot, 1.0 / 12},
		{"Convert 1in to yd", 1, LengthUnitInch, LengthUnitYard, 1.0 / 36},
		{"Convert 1in to mi", 1, LengthUnitInch, LengthUnitMile, 1.0 / 63360},

		// Foot
		{"Convert 1ft to mm", 1, LengthUnitFoot, LengthUnitMillimeter, 304.8},
		{"Convert 1ft to cm", 1, LengthUnitFoot, LengthUnitCentimeter, 30.48},
		{"Convert 1ft to m", 1, LengthUnitFoot, LengthUnitMeter, 0.3048},
		{"Convert 1ft to km", 1, LengthUnitFoot, LengthUnitKilometer, 0.0003048},
		{"Convert 1ft to in", 1, LengthUnitFoot, LengthUnitInch, 12},
		{"Convert 1ft to yd", 1, LengthUnitFoot, LengthUnitYard, 1.0 / 3},
		{"Convert 1ft to mi", 1, LengthUnitFoot, LengthUnitMile, 1.0 / 5280},

		// Yard
		{"Convert 1yd to mm", 1, LengthUnitYard, LengthUnitMillimeter, 914.4},
		{"Convert 1yd to cm", 1, LengthUnitYard, LengthUnitCentimeter, 91.44},
		{"Convert 1yd to m", 1, LengthUnitYard, LengthUnitMeter, 0.9144},
		{"Convert 1yd to km", 1, LengthUnitYard, LengthUnitKilometer, 0.0009144},
		{"Convert 1yd to in", 1, LengthUnitYard, LengthUnitInch, 36},
		{"Convert 1yd to ft", 1, LengthUnitYard, LengthUnitFoot, 3},
		{"Convert 1yd to mi", 1, LengthUnitYard, LengthUnitMile, 1.0 / 1760},

		// Mile
		{"Convert 1mi to mm", 1, LengthUnitMile, LengthUnitMillimeter, 1609344},
		{"Convert 1mi to cm", 1, LengthUnitMile, LengthUnitCentimeter, 160934.4},
		{"Convert 1mi to m", 1, LengthUnitMile, LengthUnitMeter, 1609.344},
		{"Convert 1mi to km", 1, LengthUnitMile, LengthUnitKilometer, 1.609344},
		{"Convert 1mi to in", 1, LengthUnitMile, LengthUnitInch, 63360},
		{"Convert 1mi to ft", 1, LengthUnitMile, LengthUnitFoot, 5280},
		{"Convert 1mi to yd", 1, LengthUnitMile, LengthUnitYard, 1760},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := converter.Convert(tt.value, tt.unitFrom, tt.unitTo)
			if fmt.Sprintf("%.5f", got) != fmt.Sprintf("%.5f", tt.want) {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
