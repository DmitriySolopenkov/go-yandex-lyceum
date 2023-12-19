package main

import "fmt"

type Cat struct {
	Name string
}

// Конструктор
func NewCat(name string) Cat {
	return Cat{Name: name}
}

func (c Cat) Meow() {
	fmt.Printf("%s мяукнул(а)!\n", c.Name)
}

func main() {
	myCat := NewCat("Барсик")
	myCat.Meow()
}
