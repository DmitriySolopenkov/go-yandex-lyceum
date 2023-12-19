package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "John", Age: 30}
	v := reflect.ValueOf(&p).Elem()
	v.FieldByName("Name").SetString("Jane")
	v.FieldByName("Age").SetInt(25)
	fmt.Println(p)
}
