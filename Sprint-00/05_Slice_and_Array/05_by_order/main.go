package main

import "fmt"

func SortSlice(slice []int) []int {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			slice[i-1], slice[i] = slice[i], slice[i-1]
			return SortSlice(slice)
		}
	}
	return slice
}

func main() {
	fmt.Println(SortSlice([]int{2, -4, 3, 3, 2, 1, -3})) //[-4 -3 2]
}
