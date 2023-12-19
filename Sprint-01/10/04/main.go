package main

import "fmt"

func Filter[T any](arr []T, predicate func(T) bool) []T {
	var res []T
	for _, v := range arr {
		if predicate(v) {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	result := Filter(arr, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println(result) // Output: [2 4]

	// Еще один пример:

	arr1 := []int{1, 2, 3, 4, 5}
	result1 := Filter(arr1, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println(result1) // Output: [2 4]
}

/*
Дженерик Фильтр
Создайте универсальную функцию Filter[T any](arr []T, predicate func(T) bool) []T,
	которая фильтрует элементы слайса на основе заданной функци.

Функция должна принимать срез и функции, возвращая новый срез,
	содержащий только те элементы, которые удовлетворяют предикату.

Примечания
Пример работы функции

arr := []int{1, 2, 3, 4, 5}
result := Filter(arr, func(x int) bool {
    return x%2 == 0
})
fmt.Println(result) // Output: [2 4]
Еще один пример:

arr := []int{1, 2, 3, 4, 5}
result := Filter(arr, func(x int) bool {
    return x%2 == 0
})
fmt.Println(result) // Output: [2 4]
*/
