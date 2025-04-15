package main

import (
	"log"
	"net/http"

	"github.com/spalqui/unit-converter/handlers"
)

func main() {
	homeHandler := handlers.NewHomeHandler()

	lengthHandler := handlers.NewLengthHandler()
	weightHandler := handlers.NewWeightHandler()
	temperatureHandler := handlers.NewTemperatureHandler()

	http.HandleFunc("GET /", homeHandler.Get)
	http.HandleFunc("GET /length", lengthHandler.Get)
	http.HandleFunc("GET /weight", weightHandler.Get)
	http.HandleFunc("GET /temperature", temperatureHandler.Get)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
