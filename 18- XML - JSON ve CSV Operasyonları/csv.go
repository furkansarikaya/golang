package main

import (
	"encoding/csv"
	"log"
	"os"
)

// CSV : Comma-Separated Values

func main() {
	f, err := os.Open("data.csv")
	checkError(err)

	r := csv.NewReader(f)
	r.Read()

	records, err := r.ReadAll()
	checkError(err)

	for _, row := range records {
		printRow(row)
	}
}

func printRow(row []string) {
	log.Printf("len(row) %d\n", len(row))

	for i, col := range row {
		log.Printf("[%d] : %s\n", i, col)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
