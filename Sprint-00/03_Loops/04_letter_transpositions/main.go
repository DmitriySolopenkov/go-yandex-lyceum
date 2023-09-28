package main

import "fmt"

func main() {
	var a int
	res := 1
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Некорректный ввод")
	} else {
		for i := 1; i <= a; i++ {
			res *= i
		}
		fmt.Println(res)
	}
}
