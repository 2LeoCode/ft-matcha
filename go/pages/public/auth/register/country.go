package register

import (
	"errors"
	"matcha/go/globals"
	"net/http"
	"strconv"
)

type countryForm struct {
	country int64
}

func (this *countryForm) parse(r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return errors.New("Invalid form data")
	}
	if input := r.Form.Get("country"); input == "" || input == "null" {
		return errors.New("Missing or null country in form data")
	} else if country, err := strconv.ParseInt(input, 10, 0); err != nil || country < 0 || country >= int64(len(globals.Countries)) {
		return errors.New("Invalid value for country in form data")
	} else {
		this.country = country
	}
	return nil
}

var countryHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var form countryForm

	if err := form.parse(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var start, end int
findCities:
	for i, city := range globals.Cities {
		if city.Country == form.country {
			start = i
			for j, city := range globals.Cities[i:] {
				if city.Country != form.country {
					end = j
					break findCities
				}
			}
		}
	}

	cities(start, end).Render(r.Context(), w)
})
