package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func ProcessCSV(w io.Writer) {
	file, err := os.Create("file.csv")
	if err != nil {
		log.Fatal("Error")
	}
	defer file.Close()

	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}
	wr := csv.NewWriter(file)
	for _, record := range records {
		if err := wr.Write(record); err != nil {
			log.Fatalln("Error writing record to csv: ", err)
		}
	}
	wr.Flush()
	if err := wr.Error(); err != nil {
		log.Fatal(err)
	}
}