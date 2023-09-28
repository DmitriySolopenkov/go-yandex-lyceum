package main

import "fmt"

func CalculateSeriesSum(n int) float64 {
	var res float64 = 1
	for i := 2; i <= n; i++ {
		f := float64(i)
		res += (1.0 / f)
	}
	return res
}

func main() {
	fmt.Println(CalculateSeriesSum(3))
}
