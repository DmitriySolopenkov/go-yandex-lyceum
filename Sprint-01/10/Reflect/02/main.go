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
	v := reflect.ValueOf(p)
	name := v.FieldByName("Name").String()
	age := v.FieldByName("Age").Int()
	fmt.Println(name, age)
}
