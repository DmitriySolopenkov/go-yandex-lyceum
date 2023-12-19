package main

import (
	"fmt"
	"reflect"
)

func myFunc(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Printf("Type of '%v' is %v\n", a, t)
}

func main() {
	myFunc("hello")
	myFunc(42)
}
