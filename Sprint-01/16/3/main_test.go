package main

import (
	"encoding/csv"
	"os"
	"testing"
)

func createCSVFile(filePath string, content [][]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	for _, record := range content {
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	writer.Flush()
	return writer.Error()
}

func TestSumUp(t *testing.T) {
	testCases := []struct {
		FileContent   [][]string
		ColName       string
		Expected      int
		ExpectedError bool
	}{
		{
			FileContent: [][]string{
				{"Value", "Category"},
				{"5", "A"},
				{"10", "B"},
			},
			ColName:  "Value",
			Expected: 15,
		},
		{
			FileContent:   [][]string{},
			ColName:       "Amount",
			ExpectedError: true,
		},
		{
			FileContent: [][]string{
				{"Value", "Category"},
				{"5", "A"},
				{"10", "B"},
			},
			ColName:       "Amount",
			ExpectedError: true,
		},
	}
}
