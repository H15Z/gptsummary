package csvloader

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/H15Z/gptsummary/domain/models"
)

type CSVLoader struct{}

func NewCSVLoader() *CSVLoader {
	return &CSVLoader{}
}

func (c CSVLoader) StreamData(callback func(models.Article)) {
	LoadCSV(callback)
}

func LoadCSV(callback func(models.Article)) {
	// Open the CSV file
	file, err := os.Open("data/medium_articles.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all CSV records
	header_skipped := false

	for {
		record, err := reader.Read()
		if err != nil {
			break // Reached the end of the file
		}

		if !header_skipped {
			header_skipped = true
			continue
		}

		row := models.Article{
			Title: record[0],
			Text:  record[1],
		}

		callback(row)

	}

}
