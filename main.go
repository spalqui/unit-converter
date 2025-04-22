package main

import (
	"log"
	"net/http"

	"github.com/spalqui/unit-converter/handlers/length"
	"github.com/spalqui/unit-converter/handlers/temperature"
	"github.com/spalqui/unit-converter/handlers/weight"
	"github.com/spalqui/unit-converter/services/converter"
)

func main() {
	unitConverter := converter.NewUnitConverter()

	lengthHandler := length.NewHandler(unitConverter)
	weightHandler := weight.NewHandler(unitConverter)
	temperatureHandler := temperature.NewHandler(unitConverter)

	routes := map[string]func(http.ResponseWriter, *http.Request){
		"GET /":             lengthHandler.Get,
		"GET /length":       lengthHandler.Get,
		"POST /length":      lengthHandler.Post,
		"GET /weight":       weightHandler.Get,
		"POST /weight":      weightHandler.Post,
		"GET /temperature":  temperatureHandler.Get,
		"POST /temperature": temperatureHandler.Post,
	}

	for path, handler := range routes {
		http.HandleFunc(path, handler)
	}

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
