package main

import "fmt"

type Employee struct {
	name     string
	position string
	salary   int
	bonus    int
}

func (e Employee) CalculateTotalSalary() int {
	return e.salary + e.bonus
}

func main() {
	// person1 := Employee{name: "Гоша", position: "Инженер", salary: 10, bonus: 5}
	person1 := Employee{salary: 10, bonus: 5}
	fmt.Println(person1)

}
