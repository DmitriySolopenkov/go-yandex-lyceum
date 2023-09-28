package main

import "fmt"

func MoveZeroes(nums []int) []int {
	var res []int
	for _, n := range nums {
		if n != 0 {
			res = append(res, n)
		}
	}
	for len(res) < len(nums) {
		res = append(res, 0)
	}
	return res
}

func main() {
	fmt.Println(MoveZeroes([]int{0, 1, 0, 3, 12})) //
}
