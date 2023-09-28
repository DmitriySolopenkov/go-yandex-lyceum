package main

import (
	"fmt"
	"sort"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {
	// Convert strings to lowercase for case-insensitive comparison
	s1 := strings.ToLower(str1)
	s2 := strings.ToLower(str2)

	// Sort the strings
	s1Slice := strings.Split(s1, "")
	s2Slice := strings.Split(s2, "")
	sort.Strings(s1Slice)
	sort.Strings(s2Slice)

	// Check if sorted strings are equal
	return strings.Join(s1Slice, "") == strings.Join(s2Slice, "")

}

func main() {
	fmt.Println(AreAnagrams("Кабан", "банка"))
}
