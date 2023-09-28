package main

import (
	"fmt"
	"strconv"
)

func SumTwoIntegers(a, b string) (int, error) {
	i, err := strconv.Atoi(a)
	if err != nil {
		return 0, fmt.Errorf("invalid input, please provide two integers")
	}
	j, err := strconv.Atoi(b)
	if err != nil {
		return 0, fmt.Errorf("invalid input, please provide two integers")
	}
	return i + j, nil
}

func main() {
	fmt.Println(SumTwoIntegers("1", "2"))
}
