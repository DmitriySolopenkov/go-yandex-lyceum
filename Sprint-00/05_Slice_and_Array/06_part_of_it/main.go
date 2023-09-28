package main

import "fmt"

func CountSubslices(slice []int) int {
	var res, sum, mid int
	for _, s := range slice {
		sum += s
	}
	mid = sum / len(slice)
	for i := 0; i <= len(slice); i++ {
		for j := 0; j <= len(slice); j++ {
			var sum_s int
			if i <= j {
				s_s := slice[i:j]
				for _, s := range s_s {
					sum_s += s
				}
			}
			if sum_s > mid {
				res++
			}
		}
	}
	return res
}

func main() {
	fmt.Println(CountSubslices([]int{1, 2, 3}))          // 4 mid=2
	fmt.Println(CountSubslices([]int{1, 2, 3, 4, 5}))    // 11
	fmt.Println(CountSubslices([]int{1, 2, 3, 4, 5, 6})) // 11
}
