package main

import "fmt"

func SumOfValuesInMap(m map[int]int) int {
	var res int
	for _, v := range m {
		res += v
	}
	return res
}

func main() {
	fmt.Println(SumOfValuesInMap(map[int]int{10: 38, 3: 19})) // 57
}
