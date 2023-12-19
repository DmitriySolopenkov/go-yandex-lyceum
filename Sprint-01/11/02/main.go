package main

import "slices"

func SortAndMerge(left, right []int) []int {
	var result []int
	// slices.Sort(left)
	// slices.Sort(right)

	for i := 0; i < len(left); i++ {
		result = append(result, left[i])
	}
	for j := 0; j < len(right); j++ {
		result = append(result, right[j])
	}

	slices.Sort(result)
	return result
}

/*
Сортировка со слиянием

Даны два слайса.

Напишите программу, содержащую функцию SortAndMerge(left, right []int) []int,
	которая объединит слайсы в один отсортированный в два этапа:
	- отсортировать каждый слайс
	- объединить полученные слайсы в один.

Кстати, именно так работает алгоритм сортировки слиянием ( merge sort)

Примечания
Ообъединять слайсы до сортировки не допустимо.
*/
