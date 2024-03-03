package main

import "fmt"

func IsPowerOfTwoRecursive(N int) {
	for N%2 != 1 {
		N /= 2
	}
	if N == 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main() {
	IsPowerOfTwoRecursive(16)
}
