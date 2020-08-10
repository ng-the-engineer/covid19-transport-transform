package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	csvfile, err := os.Open("./csv_files/master.csv")

	if err != nil {
		log.Fatalln("Could not open the csv file", err)
	}

	rec := csv.NewReader(csvfile)

	isValid, probRows := validateFormat(rec)

	fmt.Println("Format:", isValid)
	if len(probRows) > 0 {
		fmt.Println("Problem rows:", probRows)
	}
}
