package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetValueAsFloat64(r *http.Request, key string) (float64, error) {
	value := r.FormValue(key)
	if value == "" {
		return 0, fmt.Errorf("missing value for %s", key)
	}
	return strconv.ParseFloat(value, 64)
}
