package main

import "fmt"

func main() {
	var a, b int
	_, err := fmt.Scanln(&a, &b)
	switch {
	case err != nil:
		fmt.Println("Ошибка ввода:", err)
	case a > b:
		fmt.Println("Первое число больше второго")
	case a < b:
		fmt.Println("Второе число больше первого")
	case a == b:
		fmt.Println("Числа равны")
	}

}
