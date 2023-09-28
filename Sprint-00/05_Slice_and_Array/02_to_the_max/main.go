package main

import "fmt"

func FindMinMaxInSlice(slice []int) (int, int) {
	if len(slice) == 0 {
		return 0, 0
	}
	var min, max int = slice[0], slice[0]
	for _, s := range slice {
		switch {
		case max < s:
			max = s
		case min > s:
			min = s
		}
	}

	return min, max
}

func main() {
	fmt.Println(FindMinMaxInSlice([]int{1, 5, 10, 2, -40, 3}))
	fmt.Println(FindMinMaxInSlice([]int{}))
}
