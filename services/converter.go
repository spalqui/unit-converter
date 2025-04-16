package services

type Converter interface {
	Convert()
}

type ConvertStrategy func(value float64) float64
