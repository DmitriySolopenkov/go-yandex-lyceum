package main

import (
	"fmt"
	"math"
)

func checkPrimeNumber(num int) {
	sq_root := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq_root; i++ {
		if num%i == 0 {
			fmt.Print(num, " ")
			return
			// return strconv.Itoa(num)
		}
	}
	// return "хоп"
	fmt.Print("хоп ")
	return
}

func main() {
	// var res []string
	var n int
	fmt.Scanln(&n)
	for i := 3; i <= n; i += 5 {
		checkPrimeNumber(i)
		// res = append(res, checkPrimeNumber(i))
	}
	// fmt.Println(strings.Join(res, " "))
}
