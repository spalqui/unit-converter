package handlers

import (
	"html/template"
	"net/http"
)

type Temperature struct {
	templates map[string]*template.Template
}

func NewTemperatureHandler() *Temperature {
	return &Temperature{
		templates: map[string]*template.Template{
			http.MethodGet: template.Must(template.ParseFiles("templates/layout.gohtml", "templates/pages/temperature.gohtml")),
		},
	}
}

func (h *Temperature) Get(w http.ResponseWriter, _ *http.Request) {
	err := h.templates[http.MethodGet].Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
