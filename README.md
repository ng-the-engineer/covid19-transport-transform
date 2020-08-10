# covid19-transport-transform
Transform the cover-19 transportation data from csv to json

## Unit Test

To run unit test cases, run

```
> go test
```

## Run the app
Run

```
> go run main.go validateFormat.go
```

## Data Cleaning

1. Compile the full list from [data.gov.hk](https://data.gov.hk/en-data/dataset/hk-dh-chpsebcddr-novel-infectious-agent/resource/60e220e4-eaae-4ffd-a0fa-8895e0466ede)

2. Convert the delimiter to "tab"

3. Convert the delimiter to ";" using [Online TSV Tools](https://onlinetsvtools.com/convert-tsv-to-csv)

4. Replace all the comma "," in the file content to "-"

5. Replace all the double quote "\"" in the file content to empty string

6. Run the app to validate the file.


full_list.csv -> master.csv -> master_flight_only.csv -> master_flight_only_cleaned.csv -> flights.json