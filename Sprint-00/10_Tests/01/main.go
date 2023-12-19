package main

import (
	"fmt"
	"strings"
)

func bud(a int) {
	var b []string = []string{"*"}
	if a == 2 {
		fmt.Println("* *")
	}
	for i := a; i > 1; i-- {
		b = append(b, "*")
		defer fmt.Println(strings.Join(b[:], ""))
	}
}

func stem(b int) {
	for i := 0; i < b; i++ {
		fmt.Println("*")
	}
}

func rose(a, b int) {
	bud(a)
	stem(b)
}

func main() {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	rose(a, b)
}
