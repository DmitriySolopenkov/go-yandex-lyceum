package main

import (
	"fmt"
	"strconv"
)

func ConcatStringsAndInt(str1, str2 string, num int) string {
	var res string
	res = str1 + str2 + strconv.Itoa(num)
	return res
}

func main() {
	fmt.Println(ConcatStringsAndInt("al", "himik", 239))
}
