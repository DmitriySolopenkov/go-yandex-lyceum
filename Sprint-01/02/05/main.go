package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	var result []string
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			return nil, errors.New("invalid log format")
		}

		dateStr := parts[0]
		date, err := time.Parse("02.01.2006", dateStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse date: %v", err)
		}

		if date.After(end) {
			break
		}

		if date.After(start) || date.Equal(start) {
			result = append(result, line)
		}
	}

	if len(result) == 0 {
		return nil, errors.New("no logs found in the specified range")
	}

	return result, nil
}

func main() {

}

/*
Представьте, что другая программа пишет лог-файлы,
	где каждая строка начинается с даты формата dd.MM.YYYY.

Ваша задача - написать функцию ExtractLog(inputFileName string, start, end time.Time) ([]string, error),
	которая вернет строки "лога", созданные в указанный диапазон времени [start..end]

Формат вывода
Если в процессе работы возникает любая ошибка, то необходимо вернуть nil, err.

Если ошибок нет, то возвращайте nil в качестве значения ошибки.

Если ни одна строка не попала в указанный диапазон, то должна также вернуться произвольная ошибка.

Примечания
Например, для исходного файла:
12.12.2022 info
13.12.2022 info
14.12.2022 info
15.12.2022 info
16.12.2022 info
Если start = 13.12.2022, end = 15.12.2022, то функция должна вернуть:
13.12.2022 info
14.12.2022 info
15.12.2022 info

Если ни одна строка не попала в указанный диапазон, то должна вернуться ошибка.


*/
