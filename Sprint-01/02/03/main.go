package main

import (
	"os"
)

func ReadContent(filename string) string {
	f, err := os.ReadFile(filename)
	if err != nil {
		return string("")
	}
	return string(f)
}

func main() {
	// fmt.Println(ReadContent("file1.txt"))
}

/*
Напишите функцию ReadContent(filename string) string,
	которая принимает на вход путь к файлу,
	а возвращает его содержимое.

В случае любой ошибки возвращайте пустую строку.

Примечания
Функцию main описывать не требуется
*/
