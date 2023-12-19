package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// SumUp reads a CSV file and sums up the values in the specified column.
func SumUp(filePath, colName string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return 0, fmt.Errorf("failed to read csv: %v", err)
	}

	if len(records) < 2 { // There should be at least one data row plus the header
		return 0, errors.New("csv file does not contain enough rows")
	}

	colIndex := -1
	for i, header := range records[0] {
		if header == colName {
			colIndex = i
			break
		}
	}

	if colIndex == -1 {
		return 0, fmt.Errorf("column '%v' not found", colName)
	}

	sum := 0
	for _, record := range records[1:] {
		value, err := strconv.Atoi(record[colIndex])
		if err != nil {
			return 0, fmt.Errorf("error converting '%v' to int: %v", record[colIndex], err)
		}
		sum += value
	}

	return sum, nil
}

func main() {
	filepath := "data.csv"
	colname := "some_column"
	sum, err := SumUp(filepath, colname)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Sum: %d", sum)
}

/*
CSV
CSV (от англ. Comma-Separated Values — значения, разделённые запятыми) — текстовый формат,
	предназначенный для представления табличных данных.

Строка таблицы соответствует строке текста, которая содержит одно или несколько полей, разделенных запятыми.

Напишите функцию SumUp(filepath, colname string) (int, error), которая читает файл формата csv и суммирует значения из колонки colname. Верните полученную сумму если нет ошибок, иначе 0 и ошибку.

Примечания
первая строка файла - имена колонок.
*/
