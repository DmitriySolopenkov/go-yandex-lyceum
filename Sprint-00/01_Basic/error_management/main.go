package main

import (
	"fmt"
	"time"
)

func main() {
	var age int
	fmt.Println("Введите ваш возраст:")
	_, err := fmt.Scanln(&age)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(age)
	}
	now := time.Now() // вызываем команду, которая кладёт в переменную текущее время
	fmt.Println(now)
}
