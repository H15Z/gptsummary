package csvloader

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/H15Z/gptsummary/domain/models"
)

func LoadCSV() {
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
	var rows []models.Article
	for {
		record, err := reader.Read()
		if err != nil {
			break // Reached the end of the file
		}

		//TODO TRIGGER CALLBACK METHOD EACH TIME A ROW IS PROCESSED
		row := models.Article{
			Title: record[0],
			Text:  record[1],
		}

		rows = append(rows, row)
	}

	// Print the parsed data
	for _, row := range rows {
		fmt.Printf("Title: %s\n", row.Title)
		fmt.Printf("Text: %s\n", row.Text)
		fmt.Println("---------")
	}
}
