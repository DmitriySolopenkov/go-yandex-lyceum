package main

import (
	"fmt"
	"math"
)

func FindIntersection(k1, b1, k2, b2 float64) (float64, float64) {
	var x, y float64
	if k1 != k2 {
		x = (b2 - b1) / (k1 - k2)
		y = (k1 * x) + b1
	} else {
		x, y = math.NaN(), math.NaN()
	}
	return x, y
}

func main() {
	fmt.Println(FindIntersection(5, 20, -3, 10)) // -1.25 13.75
}
