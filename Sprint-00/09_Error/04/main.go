package main

import "fmt"

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial is not defined for negative numbers")
	}
	// крайний случай
	if n == 0 {
		return 1, nil
	}
	// вызов функции с текущим значением умноженным на результат (n-1)
	total := 1
	for i := 1; i <= n; i++ {
		total *= i
	}
	return total, nil
}

func main() {
	fmt.Println(Factorial(5))
}
