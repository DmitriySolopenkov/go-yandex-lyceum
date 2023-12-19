package main

func Clean(nums []int, x int) []int {
	i := 0
	for _, val := range nums {
		if val != x {
			nums[i] = val
			i++
		}
	}
	return nums[:i]
}

/*
Удаление элемента


Дан неотсортированный слайс целых чисел.
	Напишите функцию Clean(nums []int, x int) ([]int),
	которая удаляет из исходного слайса все вхождения x.
	Важно - не использовать дополнительный слайс.

Примечания
Функцию main создавать не надо.
*/
