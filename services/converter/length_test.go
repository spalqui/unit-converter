package converter

import (
	"fmt"
	"testing"
)

func Test_millimeterToCentimeter(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{
			name:  "Convert 10mm to cm",
			value: 10,
			want:  1,
		},
		{
			name:  "Convert 0mm to cm",
			value: 0,
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := millimeterToCentimeter(tt.value); got != tt.want {
				t.Errorf("millimeterToCentimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_millimeterToMeter(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{
			name:  "Convert 1000mm to m",
			value: 1000,
			want:  1,
		},
		{
			name:  "Convert 0mm to m",
			value: 0,
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := millimeterToMeter(tt.value); got != tt.want {
				t.Errorf("millimeterToMeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_millimeterToKilometer(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{
			name:  "Convert 1,000,000mm to km",
			value: 1_000_000,
			want:  1,
		},
		{
			name:  "Convert 0mm to km",
			value: 0,
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := millimeterToKilometer(tt.value); got != tt.want {
				t.Errorf("millimeterToKilometer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_millimeterToInch(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{
			name:  "Convert 100mm to in",
			value: 100,
			want:  3.93701,
		},
		{
			name:  "Convert 0mm to in",
			value: 0,
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := millimeterToInch(tt.value); fmt.Sprintf("%.5f", got) != fmt.Sprintf("%.5f", tt.want) {
				t.Errorf("millimeterToInch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_millimeterToFoot(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{
			name:  "Convert 1700mm to ft",
			value: 1700,
			want:  5.57743,
		},
		{
			name:  "Convert 0mm to in",
			value: 0,
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := millimeterToFoot(tt.value); fmt.Sprintf("%.5f", got) != fmt.Sprintf("%.5f", tt.want) {
				t.Errorf("millimeterToFoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_millimeterToYard(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{
			name:  "Convert 1500mm to yd",
			value: 1500,
			want:  1.64042,
		},
		{
			name:  "Convert 0mm to yd",
			value: 0,
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := millimeterToYard(tt.value); fmt.Sprintf("%.5f", got) != fmt.Sprintf("%.5f", tt.want) {
				t.Errorf("millimeterToYard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_millimeterToMile(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{
			name:  "Convert 804672mm to mi",
			value: 804672,
			want:  0.5,
		},
		{
			name:  "Convert 0mm to mi",
			value: 0,
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := millimeterToMile(tt.value); fmt.Sprintf("%.5f", got) != fmt.Sprintf("%.5f", tt.want) {
				t.Errorf("millimeterToMile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_centimeterToMillimeter(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{
			name:  "Convert 1cm to mm",
			value: 1,
			want:  10,
		},
		{
			name:  "Convert 0cm to mm",
			value: 0,
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := centimeterToMillimeter(tt.value); got != tt.want {
				t.Errorf("centimeterToMillimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
