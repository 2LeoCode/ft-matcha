package globals

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type City struct {
	Name      string
	Country   int64
	Latitude  float64
	Longitude float64
}

var Cities []City

func init() {
	var (
		err        error
		f          *os.File
		csvEntries [][]string
	)

	if f, err = os.Open("assets/cities.csv"); err == nil {
		csvEntries, err = csv.NewReader(f).ReadAll()
	}
	if err != nil {
		log.Fatalf("Error while loading assets/countries.csv: %s", err.Error())
	}
	Cities = make([]City, len(csvEntries))
	for i, row := range csvEntries {
		var _ error

		Cities[i].Name = row[0]
		Cities[i].Country, _ = strconv.ParseInt(row[1], 0, 0)
		Cities[i].Latitude, _ = strconv.ParseFloat(row[2], 0)
		Cities[i].Longitude, _ = strconv.ParseFloat(row[3], 0)
	}
}
