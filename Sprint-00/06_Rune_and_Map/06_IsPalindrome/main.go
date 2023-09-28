package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(input string) bool {
	str := strings.Join(strings.Fields(strings.ToLower(input)), "")
	l := len([]rune(str))
	for i := 0; i < l/2; i++ {
		if []rune(str)[i] != []rune(str)[l-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome("Ароза упала на лапу Азора")) //true
}
