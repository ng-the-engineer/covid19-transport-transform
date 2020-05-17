package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	csvfile, err := os.Open("covid19.csv")

	if err != nil {
		log.Fatalln("Could not open the csv file", err)
	}

	rec := csv.NewReader(csvfile)

	isValid, probRows := validateFormat(rec)

	fmt.Println("Row(s) has problem:", isValid, probRows)
}
