package handlers

import (
	"html/template"
	"net/http"

	"github.com/spalqui/unit-converter/templates/pages"
)

type Length struct {
	templates map[string]*template.Template
}

func NewLengthHandler() *Length {
	return &Length{
		templates: map[string]*template.Template{
			http.MethodGet: template.Must(template.ParseFiles("templates/layout.gohtml", "templates/pages/length.gohtml")),
		},
	}
}

func (h *Length) Get(w http.ResponseWriter, _ *http.Request) {
	page := pages.Page{
		Title: "Length",
		Units: []string{
			"Millimeter",
			"Centimeter",
			"Meter",
			"Kilometer",
			"Inch",
			"Foot",
			"Yard",
			"Mile",
		},
	}

	err := h.templates[http.MethodGet].Execute(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
