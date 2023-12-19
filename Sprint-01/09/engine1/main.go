package main

import "fmt"

type Engine struct {
	Model string
}

type Car struct {
	Make string
	Eng  Engine
}

func main() {
	engine := Engine{Model: "V6"}
	car := Car{Make: "Toyota", Eng: engine}

	fmt.Printf("Car Make: %s\n", car.Make)
	fmt.Printf("Engine Model: %s\n", car.Eng.Model)
}
