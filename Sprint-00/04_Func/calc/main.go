package main

import (
	"fmt"
)

func sum(a, b int) int {
	return a + b
}

func dedc(a, b int) int {
	return a - b
}

func div(a, b int) int {
	return a * b
}

func main() {
	var i, j int
	fmt.Println("Введите два числа:")
	_, err := fmt.Scanln(&i, &j)
	if err != nil {
		fmt.Println("Некорректный ввод!")
	} else {
		fmt.Println("Сумма равна:", sum(i, j))
		fmt.Println("Разность равна:", dedc(i, j))
		fmt.Println("Произведение равно:", div(i, j))
	}

}
