package main

import "fmt"

func main() {
	var a int
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Некорректный ввод")
	} else {
		for i := 1; i <= 10; i++ {
			fmt.Println(a, "*", i, "=", a*i)
		}
	}
}
