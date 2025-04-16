package handlers

import (
	"html/template"
	"net/http"
)

type Weight struct {
	templates map[string]*template.Template
}

func NewWeightHandler() *Weight {
	return &Weight{
		templates: map[string]*template.Template{
			http.MethodGet: template.Must(template.ParseFiles("templates/layout.gohtml", "templates/pages/weight.gohtml")),
		},
	}
}

func (h *Weight) Get(w http.ResponseWriter, _ *http.Request) {
	err := h.templates[http.MethodGet].Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
