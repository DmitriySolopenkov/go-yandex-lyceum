package LineByNum

import (
	"bufio"
	"os"
)

func LineByNum(inputFilename string, lineNum int) string {
	f, err := os.Open(inputFilename)
	if err != nil {
		return "Error"
	}
	defer f.Close()
	l := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if l == lineNum {
			return string(scanner.Text())
		}
		l++
	}
	return ""
}

/*
Напишите функцию LineByNum(inputFilename string, lineNum int) string,
	которая получает в качестве параметров имя файла и номер строи,
	а возвращает текст строки по ее порядковому номеру в файле.

Если строки с указанным номером найти не удается, то верните пустую строку.

Примечания
Функцию main создавать не надо.
Нумерация элементов в программировании всегда начинается с 0.


*/
