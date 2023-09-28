package main

import "fmt"

func ReverseSlice(slice []int) []int {
	var res []int
	len := len(slice)
	for i := (len - 1); i >= 0; i-- {
		// s := slice[i]
		res = append(res, slice[i])
	}
	return res
}

func main() {
	fmt.Println(ReverseSlice([]int{1, 2, 3, 4, 5}))
}
