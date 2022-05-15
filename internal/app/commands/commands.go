package commands

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func readCsv(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	// csvReader.Comma = ';'

	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// convert records to array of structs
	// ffList := createFfList(data)
	// return ffList, err
	return data, err
}

func addRow(data [][]string, row []string) ([][]string, error) {
	fmt.Println(row)
	data = append(data, row)
	fmt.Println(data)

	f, err := os.Create("test.csv")
	defer f.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(f)
	err = w.WriteAll(data)

	if err != nil {
		log.Fatal(err)
	}
	return data, err
}
