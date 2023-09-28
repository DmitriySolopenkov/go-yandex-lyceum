package main

import (
	"fmt"
	"math"
)

func SqRoots() {
	var a, b, c float64
	_, err := fmt.Scanln(&a, &b, &c)
	if err != nil || a == 0 {
		fmt.Println("Некорректный ввод!")
	} else {
		D := b*b - 4.0*a*c
		switch {
		case D < 0.0:
			fmt.Println("0", "0")
		case D > 0.0:
			fmt.Println(((-b - math.Sqrt(D)) / (2 * a)), ((-b + math.Sqrt(D)) / (2 * a)))
		case D == 0.0:
			fmt.Println((-b / (2 * a)))
		}
	}
}

func main() {
	SqRoots()
}

/*
ax2 + bx + c = 0
D = b2 - 4ac
При D > 0, уравнение имеет два корня;
при D = 0, один корень;
при D < 0, уравнение корней не имеет

math.Sqrt(D)

*/
