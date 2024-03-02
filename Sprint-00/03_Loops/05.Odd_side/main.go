package main

import "fmt"

func main() {
	var a, res int
	_, err := fmt.Scanln(&a)
	if err != nil || a < 0 {
		fmt.Println("Некорректный ввод")
	} else {
		for i := 1; i <= a; i++ {
			if i%2 == 0 {
				continue
			} else {
				res += i
			}
		}
		fmt.Println(res)
	}
}
