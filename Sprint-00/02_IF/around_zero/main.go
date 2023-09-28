package main

import "fmt"

func main() {
	var a, b int
	_, err := fmt.Scanln(&a, &b)
	switch {
	case err != nil:
		fmt.Println("Ошибка ввода:", err)
	case a == 0 || b == 0:
		fmt.Println("Одно из чисел равно нулю")
	case a > 0 && b > 0:
		fmt.Println("Оба числа положительные")
	case a < 0 && b < 0:
		fmt.Println("Оба числа отрицательные")
	default:
		fmt.Println("Одно число положительное, а другое отрицательное")
	}
}
