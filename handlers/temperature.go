package handlers

import (
	"html/template"
	"net/http"

	"github.com/spalqui/unit-converter/templates/pages"
)

type Temperature struct {
	templates map[string]*template.Template
}

func NewTemperatureHandler() *Temperature {
	return &Temperature{
		templates: map[string]*template.Template{
			http.MethodGet: template.Must(template.ParseFiles("templates/layout.gohtml", "templates/pages/converter.gohtml")),
		},
	}
}

func (h *Temperature) Get(w http.ResponseWriter, _ *http.Request) {
	page := pages.Converter{
		Title: "Temperature",
	}

	err := h.templates[http.MethodGet].Execute(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
