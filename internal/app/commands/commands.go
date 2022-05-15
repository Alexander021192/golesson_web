package commands

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Art struct {
	Id       string
	Title    string
	Anons    string
	FullText string
}

func ReadCsv(filename string) ([][]string, error) {
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

func AddRow(data [][]string, row []string) ([][]string, error) {
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

func CreateArtList(data [][]string) []Art {
	var ArtList []Art
	for i, line := range data {
		if i > 0 { // omit header line
			var row Art
			for j, field := range line {
				switch j {
				case 0:
					row.Id = field
				case 1:
					row.Title = field
				case 2:
					row.Anons = field
				case 3:
					row.FullText = field
				}
			}
			ArtList = append(ArtList, row)
		}
	}
	return ArtList
}