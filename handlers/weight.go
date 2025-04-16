package handlers

import (
	"html/template"
	"net/http"

	"github.com/spalqui/unit-converter/templates/pages"
)

type Weight struct {
	templates map[string]*template.Template
}

func NewWeightHandler() *Weight {
	return &Weight{
		templates: map[string]*template.Template{
			http.MethodGet: template.Must(template.ParseFiles("templates/layout.gohtml", "templates/pages/converter.gohtml")),
		},
	}
}

func (h *Weight) Get(w http.ResponseWriter, _ *http.Request) {
	page := pages.Converter{
		Title: "Weight",
	}

	err := h.templates[http.MethodGet].Execute(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
