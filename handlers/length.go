package handlers

import (
	"html/template"
	"net/http"

	"github.com/spalqui/unit-converter/services/converter"
	"github.com/spalqui/unit-converter/templates/pages"
)

type Length struct {
	templates map[string]*template.Template
	service   *converter.LengthConverter
}

func NewLengthHandler(service *converter.LengthConverter) *Length {
	return &Length{
		templates: map[string]*template.Template{
			http.MethodGet: template.Must(template.ParseFiles("templates/layout.gohtml", "templates/pages/length.gohtml")),
		},
		service: service,
	}
}

func (h *Length) Get(w http.ResponseWriter, _ *http.Request) {
	page := pages.Page{
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

func (h *Length) Post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value, err := getValueAsFloat64(r, "value")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	unitFrom := r.FormValue("unitFrom")
	if unitFrom == "" {
		http.Error(w, "missing unit from", http.StatusBadRequest)
		return
	}

	unitTo := r.FormValue("unitTo")
	if unitTo == "" {
		http.Error(w, "missing unit to", http.StatusBadRequest)
		return
	}

	result := h.service.Convert(value, unitFrom, unitTo)

	page := pages.Page{
		Value:    value,
		Result:   result,
		UnitFrom: r.FormValue("unitFrom"),
		UnitTo:   r.FormValue("unitTo"),
	}

	err = h.templates[http.MethodGet].Execute(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
