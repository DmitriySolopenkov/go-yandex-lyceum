package main

import (
	"fmt"
	"strconv"
)

func IntToBinary(num int) (string, error) {
	if num < 0 {
		return "", fmt.Errorf("negative numbers are not allowed")
	}
	return strconv.FormatInt(int64(num), 2), nil
}

func main() {
	fmt.Println(IntToBinary(-1))
}

//negative numbers are not allowed
