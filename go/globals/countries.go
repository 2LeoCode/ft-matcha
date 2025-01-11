package globals

import (
	"encoding/csv"
	"log"
	"os"
)

var Countries []string

func init() {
	var (
		err        error
		f          *os.File
		csvEntries [][]string
	)

	if f, err = os.Open("assets/countries.csv"); err == nil {
		csvEntries, err = csv.NewReader(f).ReadAll()
	}
	if err != nil {
		log.Fatalf("Error while loading assets/countries.csv: %s", err.Error())
	}
	Countries = make([]string, len(csvEntries))
	for i, row := range csvEntries {
		Countries[i] = row[0]
	}
}
