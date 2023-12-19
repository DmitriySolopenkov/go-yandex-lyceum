package main

func SliceCopy(nums []int) []int {
	var result []int
	result = append(result, nums...)
	return result
}

/*
Копирование


Дан слайс целых чисел nums.
	Этот слайс имеет емкость больше его длины.
	Создайте функцию SliceCopy(nums []int) []int,
	которая вернёт новый слайс длиной и ёмкостью,
	равной длине nums.
	Скопируйте в него значения из исходного слайса.

Примечания
Функцию main создавать не надо
*/
