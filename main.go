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

	http.HandleFunc("GET /", lengthHandler.Get)

	http.HandleFunc("GET /length", lengthHandler.Get)
	http.HandleFunc("POST /length", lengthHandler.Post)

	http.HandleFunc("GET /weight", weightHandler.Get)
	http.HandleFunc("POST /weight", weightHandler.Post)

	http.HandleFunc("GET /temperature", temperatureHandler.Get)
	http.HandleFunc("POST /temperature", temperatureHandler.Post)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
