package main

import "fmt"

func main() {
	var a int
	_, err := fmt.Scanln(&a)
	switch {
	case err != nil:
		fmt.Println("Ошибка ввода:", err)
	case a > 0 && a%2 == 0:
		fmt.Println("Число положительное и четное")
	case a <= 0 && a%2 == 0:
		fmt.Println("Число отрицательное и четное")
	case a > 0:
		fmt.Println("Число положительное и нечетное")
	case a <= 0:
		fmt.Println("Число отрицательное и нечетное")
	}
}
