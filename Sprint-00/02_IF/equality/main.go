package main

import "fmt"

func main() {
	var a, b, c int
	_, err := fmt.Scanln(&a, &b, &c)
	switch {
	case err != nil:
		fmt.Println("Некорректный ввод")
	case a <= 0 || b <= 0 || a <= 0:
		fmt.Println("Некорректный ввод")
	case a == b && b == c:
		fmt.Println("Все числа равны")
	case a == b || b == c || a == c:
		fmt.Println("Два числа равны")
	default:
		fmt.Println("Все числа разные")
	}
}
