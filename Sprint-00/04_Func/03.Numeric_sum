package main

import "fmt"

func SumDigitsRecursive(n int) int {
	var d []int
	var res int
	for n != 0 {
		d = append(d, n%10)
		n /= 10
	}
	for _, i := range d {
		res += i
	}
	return res
}

func main() {
	fmt.Println(SumDigitsRecursive(123))
}
