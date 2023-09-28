package main

import "fmt"

// Структуры
type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

// Методы
func (d Dog) MakeSound() {
	fmt.Println("Гав!")
}

func (c Cat) MakeSound() {
	fmt.Println("Мяу!")
}

// Интерфейс
type Animal interface {
	MakeSound()
}

func main() {
	cat1 := Cat{}
	cat1.MakeSound()
	dog1 := Dog{}
	dog1.MakeSound()
}
