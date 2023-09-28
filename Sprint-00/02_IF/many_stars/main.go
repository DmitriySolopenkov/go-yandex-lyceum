package main

import "fmt"

func main() {
	var a int
	_, err := fmt.Scanln(&a)
	switch {
	case err != nil:
		fmt.Println("Некорректный ввод")
	case a < 0:
		fmt.Println("Некорректный ввод")
	case a < 10:
		fmt.Println("Число меньше 10")
	case a < 100:
		fmt.Println("Число меньше 100")
	case a < 1000:
		fmt.Println("Число меньше 1000")
	default:
		fmt.Println("Число больше или равно 1000")
	}
}
