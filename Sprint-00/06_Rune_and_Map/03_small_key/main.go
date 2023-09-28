package main

import "fmt"

const (
	maxInt = int(^uint(0) >> 1)
	minInt = -maxInt - 1
)

func FindMaxKey(m map[int]int) int {
	var minNumber int
	for minNumber = range m {
		break
	}
	for n := range m {
		if n > minNumber {
			minNumber = n
		}
	}
	return minNumber
}

func main() {
	fmt.Println(FindMaxKey(map[int]int{-7: 1, -1: 38, 10: 37, 30: 19}))
}
