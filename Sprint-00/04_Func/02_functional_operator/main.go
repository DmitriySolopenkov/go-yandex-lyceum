package main

import "fmt"

func Add(a, b float64) float64 {
	return a + b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func PrintNumbersAscending(n int) {
	for i := 1; i <= n; i++ {
		fmt.Print(i, " ")
	}
}

func main() {
	// fmt.Println(Add(3, 7))
	// fmt.Println(Multiply(3, 7))
	PrintNumbersAscending(3)
}
