package main

import "fmt"

func Convert2D(nums []int, m, n int) [][]int {
	var res = make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = nums[i*n : (i+1)*n]
	}
	return res
}

func main() {
	fmt.Println(Convert2D([]int{1, 2, 3, 4, 5, 6, 7, 8}, 4, 2)) //[[1 2] [3 4] [5 6] [7 8]]
}
