package main

import (
	"encoding/csv"
	"log"
	"os"
	"testing"
)

func TestValidateFormatWithWrongFormat(t *testing.T) {

	expectedIsValid := false

	expectedProbRows := "3,5"

	csvfile, err := os.Open("./tests/test_wrong_format.csv")

	if err != nil {
		log.Fatalln("Could not open the csv file", err)
	}

	rec := csv.NewReader(csvfile)

	actualIsValid, actualProbRows := validateFormat(rec)

	if actualIsValid != expectedIsValid {
		t.Errorf("isValid is expected false but the actaul result is %v", actualIsValid)
	}

	if actualProbRows != expectedProbRows {
		t.Errorf("probRows is expected '3,5' but the actaul result is %v", actualProbRows)
	}
}

func TestValidateFormatWithRightFormat(t *testing.T) {
	expectedIsValid := true

	expectedProbRows := ""

	csvfile, err := os.Open("./tests/test_right_format.csv")

	if err != nil {
		log.Fatalln("Could not open the csv file", err)
	}

	rec := csv.NewReader(csvfile)

	actualIsValid, actualProbRows := validateFormat(rec)

	if actualIsValid != expectedIsValid {
		t.Errorf("isValid is expected true but the actaul result is %v", actualIsValid)
	}

	if actualProbRows != expectedProbRows {
		t.Errorf("probRows is expected empty string but the actaul result is %v", actualProbRows)
	}
}
