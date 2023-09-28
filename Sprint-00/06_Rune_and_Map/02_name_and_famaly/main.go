package main

import (
	"fmt"
	"strings"
)

func ConcatenateStrings(str1, str2 string) string {
	return strings.Join([]string{str1, str2}, " ")
}

func main() {
	fmt.Println(ConcatenateStrings("Ira", "Khomyakova"))
}
