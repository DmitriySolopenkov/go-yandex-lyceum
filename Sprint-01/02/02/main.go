package main

import (
	"fmt"
	"os"
)

func ModifyFile(filename string, pos int, val string) {
	f, err := os.OpenFile(filename, os.O_RDWR, 0600)
	if err != nil {
		fmt.Println("Error file open!")
	}
	defer f.Close()

	f.Seek(int64(pos), 0)
	f.WriteString(val)
}

func main() {

}

/*
Напишите функцию ModifyFile(filename string, pos int, val string),
	которая будет изменять содержимое файла на значение val, начиная с позиции pos.

Для перемещения по файлу используйте функцию os.Seek.

Примечания
Функцию main описывать не требуется.
*/
