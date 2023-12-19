package main

import (
	"io"
	"os"
)

func CopyFilePart(inputFilename, outFileName string, startpos int) error {
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = inputFile.Seek(int64(startpos), 0)
	if err != nil {
		return err
	}

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return err
	}

	return nil
}

func main() {

}

/*
Напишите функцию CopyFilePart(inputFilename, outFileName string, startpos int) error,
	которая открывает файл с именем inputFilename на чтение,
	создает файл с именем outFileName и записывает содержимое файла inputFilename
	с позиции startPos и до конца в файл outFileName.

Не забудьте закрыть файлы после обработки.

Заметьте, что функция возвращает ошибку. Если все операции прошли без ошибок, то верните nil
*/
