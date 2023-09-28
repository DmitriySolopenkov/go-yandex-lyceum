package main

import "fmt"

func main() {
	var input int

	fmt.Scanln(&input)

	if input > 0 {
		fmt.Println("Число положительное")
	} else {
		fmt.Println("Число отрицательное")
	}
}
