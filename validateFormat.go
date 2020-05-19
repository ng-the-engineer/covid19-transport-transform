package main

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"
)

// -----------------------------------------------------------
// Assumption
// 1. No header column
// 2. Four columns; Username; Identifier;First name;Last name
// -----------------------------------------------------------

func validateFormat(rd *csv.Reader) (bool, string) {
	rowCnt := 0

	columnCount := 5

	problemRows := []string{}

	for {
		row, err := rd.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		rowArr := strings.Split(row[0], ";")

		if len(rowArr) != columnCount {
			problemRows = append(problemRows, strconv.Itoa(rowCnt))
		}

		rowCnt++
	}

	isValid := (len(problemRows) == 0)
	output := strings.Join(problemRows, ",")

	return isValid, output
}
