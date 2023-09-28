package main

import (
	"fmt"
)

func FindValue(nums []int) int {
	var res int
	keys := make(map[int]bool)
	for _, entry := range nums {
		if _, value := keys[entry]; !value {
			keys[entry] = false
			res = entry
		}
	}
	return res
}

func main() {
	fmt.Println(FindValue([]int{2, 2, 5, 5, 1})) // 1
}
