package handlers

import (
	"html/template"
	"net/http"

	"github.com/spalqui/unit-converter/templates/pages"
)

type Home struct {
	templates map[string]*template.Template
}

func NewHomeHandler() *Home {
	return &Home{
		templates: map[string]*template.Template{
			http.MethodGet: template.Must(template.ParseFiles("templates/layout.gohtml", "templates/pages/home.gohtml")),
		},
	}
}

func (h *Home) Get(w http.ResponseWriter, _ *http.Request) {
	page := pages.Page{
		Title:  "Home",
		Result: "Hello, World!",
	}

	err := h.templates[http.MethodGet].Execute(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
