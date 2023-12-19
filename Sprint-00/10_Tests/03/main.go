package main

import (
	"fmt"
)

func main() {
	// var s1, s2 string
	s1 := "τ"
	s2 := "α"
	var c1, c2 int
	// fmt.Scanln(&s1)
	// fmt.Scanln(&s2)
	r1 := []rune(s1)
	r2 := []rune(s2)
	// in := bufio.NewReader(os.Stdin)
	// line, _ := in.ReadString('\n')
	line := "пцуατк цуτпкαц уцαу"
	for _, r := range []rune(line) {
		switch {
		case r1[0] == r:
			c1++
		case r2[0] == r:
			c2++
		}
	}
	if c1 >= c2 {
		fmt.Println(s1, c1)
		fmt.Println(s2, c2)
	} else {
		fmt.Println(s2, c2)
		fmt.Println(s1, c1)
	}

}
