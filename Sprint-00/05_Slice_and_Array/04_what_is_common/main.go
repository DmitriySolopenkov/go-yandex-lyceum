package main

import "fmt"

func IntersectionOfSlices(slice1, slice2 []int) []int {
	var res []int
	for _, s1 := range slice1 {
		for _, s2 := range slice2 {
			if s1 == s2 {
				res = append(res, s1)
			}
		}
	}
	return res
}

func main() {
	fmt.Println(IntersectionOfSlices([]int{1, 2, 3, 4}, []int{1, 2})) // [1 2]
	fmt.Println(IntersectionOfSlices([]int{1, 2, 3, 4}, []int{0}))    // []
}
