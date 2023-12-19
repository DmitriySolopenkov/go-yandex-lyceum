package main

import (
	"fmt"
	"sort"
	"strings"
)

type Goods struct {
	Names []string // Название товара
	Boxes []int    // Коробка, в которой лежит товар
}

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func (g Goods) Len() int {
	return len(g.Names)
}

func (g Goods) Swap(i, j int) {
	g.Names[i], g.Names[j] = g.Names[j], g.Names[i] // Меняем местами как название товара,
	g.Boxes[i], g.Boxes[j] = g.Boxes[j], g.Boxes[i] // так и коробки, где товар хранится
}
func (g Goods) Less(i, j int) bool {
	return g.Boxes[i] < g.Boxes[j]
}

func main() {
	intSlice := []int{4, 5, 2, 1, 3, 9, 7, 8, 6}
	fmt.Println(sort.IntsAreSorted(intSlice)) // Проверим, отсортирован ли слайс
	sort.Ints(intSlice)                       // Сама сортировка
	fmt.Println(intSlice)                     // [1 2 3 4 5 6 7 8 9]

	strs := []string{"A very long string", "a medium string", "a short one"}
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	fmt.Println(strings.Join(strs, ", ")) // a short one, a medium string, a very long string

	type User struct {
		Name string
		Age  int
	}
	users := []User{
		{
			Name: "Ivan",
			Age:  56,
		},
		{
			Name: "Tim",
			Age:  33,
		},
		{
			Name: "Bob",
			Age:  89,
		},
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})
	fmt.Println(users) // [{Tim 33} {Ivan 56} {Bob 89}]

	goods := Goods{
		Names: []string{"phone", "tablet", "earphones"},
		Boxes: []int{3, 1, 2},
	}

	sort.Sort(goods)
	fmt.Println(goods) // {[tablet earphones phone] [1 2 3]}

}
