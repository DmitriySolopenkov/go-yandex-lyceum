package main

import (
	"fmt"
	"slices"
)

func main() {
	smallInts := []int8{0, 42, -10, 8}
	slices.Sort(smallInts)
	fmt.Println(smallInts) // [-10 0 8 42]

	slices.SortFunc(smallInts, func(a, b int8) int {
		// Здесь мы сравниваем сами
		switch {
		case a < b:
			return 1
		case a > b:
			return -1
		default:
			return 0
		}
	})
	fmt.Println(smallInts) // [42 8 0 -10]

	type User struct {
		Name string
		Age  int
	}
	users := []User{
		{
			Name: "Ivan",
			Age:  56,
		},
		{
			Name: "Tim",
			Age:  33,
		},
		{
			Name: "Bob",
			Age:  89,
		},
	}
	slices.SortFunc(users, func(a, b User) int {
		// Здесь мы сравниваем сами
		switch a.Age < b.Age {
		case true:
			return -1
		case false:
			return 1
		default:
			return 0
		}
	})
	fmt.Println(users) // [{Tim 33} {Ivan 56} {Bob 89}]

}
