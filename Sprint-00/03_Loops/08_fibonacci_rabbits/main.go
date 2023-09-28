package main

import "fmt"

func main() {
	var a, j int
	_, err := fmt.Scanln(&a)
	if err != nil || a < 0 {
		fmt.Println("Некорректный ввод")
	} else {
		first, second := 0, 1
		for i := 0; i < (10 + j); i++ {
			res := first
			first, second = second, first+second
			if res < a {
				j++
				continue
			} else {
				fmt.Println(res)
			}

		}
	}
}
