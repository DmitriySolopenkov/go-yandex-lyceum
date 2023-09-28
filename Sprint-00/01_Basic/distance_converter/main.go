package main

import "fmt"

func main() {
	meters := 0.0
	fmt.Scan(&meters)

	const exchangeRate = 1852
	seaMiles := meters / exchangeRate
	fmt.Println(seaMiles)
}
