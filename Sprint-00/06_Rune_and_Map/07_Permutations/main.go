package main

import (
	"fmt"
	"sort"
)

func Permutations(input string) []string {
	var res []string
	// Сортировака строки по алфавиту
	s := []byte(input)
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	// Перестановка
	res = generatePermutationsRecursive(string(s))
	return res
}

func generatePermutationsRecursive(str string) []string {
	if len(str) == 1 {
		return []string{str}
	}
	permutations := []string{}
	for i, c := range str {
		remaining := str[:i] + str[i+1:]
		subPermutations := generatePermutationsRecursive(remaining)
		for _, p := range subPermutations {
			permutations = append(permutations, string(c)+p)
		}
	}
	return permutations
}

func main() {
	fmt.Println(Permutations("bac")) //[abc acb bac bca cab cba]
}
