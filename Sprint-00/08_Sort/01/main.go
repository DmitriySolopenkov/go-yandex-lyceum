package main

import (
	"fmt"
)

func FindMissingValues(nums []int) []int {
	var res []int
	len := len(nums)
	seen := make([]bool, len+1)
	// Проверка наличия чисел в слайсе
	for i := 0; i < len; i++ {
		seen[nums[i]] = true
	}
	// Создание нового класа, для елементов не найденных в слайсе "nums"
	for i := 1; i <= len; i++ {
		if !seen[i] {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	fmt.Println(FindMissingValues([]int{4, 3, 2, 7, 8, 2, 3, 1})) //  [5 6]
	fmt.Println(FindMissingValues([]int{4, 4, 4, 4}))             //  [1 2 3]
	fmt.Println(FindMissingValues([]int{1, 1, 1}))                //  [2 3]

}
