package main

import "fmt"

func main() {
	var a, res int
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Некорректный ввод")
	} else {
		for i := 1; i <= a; i++ {
			res += i
		}
		fmt.Println(res)
	}
}
