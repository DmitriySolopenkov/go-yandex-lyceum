package main

import "fmt"

func SumOfSlice(slice []int) int {
	var res int
	for _, s := range slice {
		res += s
	}

	return res
}

func main() {
	fmt.Println(SumOfSlice([]int{1, 2, 3, 4})) // 10
}
