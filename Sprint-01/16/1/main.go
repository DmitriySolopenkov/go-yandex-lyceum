package main

func Rotate(data []int, pos int) []int {
	var res []int
	switch {
	case pos == 0:
		res = data
	case pos > 0:
		pr := (pos % (len(data))) + 1
		res = append(data[pr:], data[:pr]...)
	case pos < 0:
		pr := (pos % (len(data) - 1))
		pr += len(data) - 1
		res = append(data[pr:], data[:pr]...)
	}

	return res
}

/*
Сдвиг

Напишите функцию Rotate(data []int, pos int) []int,
	которая осуществляет циклический сдвиг элементов слайса чисел на заданное количество позиций.

	Пример: если data = [1,2,3,4,5,6,7] и pos = 3, то функция должна вернуть [5,6,7,1,2,3,4]

Примечания
pos может быть отрицательным числом, а также больше длины слайса
*/
