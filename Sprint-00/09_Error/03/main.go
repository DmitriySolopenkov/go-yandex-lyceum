package main

import (
	"errors"
	"fmt"
)

func GetCharacterAtPosition(str string, position int) (rune, error) {
	r := []rune(str)
	if position < 0 || position >= len(r) {
		return 0, errors.New("position out of range")
	}
	return r[position], nil
}

func main() {
	fmt.Println(GetCharacterAtPosition("Зенит чемпион!", 20))

}
