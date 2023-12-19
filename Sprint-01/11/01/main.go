package main

import "sort"

func SortNums(nums []uint) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
}

/*
Сортировка слайса uint
Напишите программу, содержащую функцию SortNums(nums []uint), которая сортирует слайс nums по возрастанию
*/
