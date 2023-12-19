package main

func Sum[T Number](arr []T) T {
	var total T
	for _, n := range arr {
		total += n
	}
	return total
}

type Number interface {
	int | float64
}

/*
Суммирование
Напишите универсальную функцию Sum[T Number](arr []T) T,
	которая может суммировать элементы срезов разных числовых типов (например, int, float64).

Функция должна брать фрагмент любого числового типа и возвращать его сумму.

Примечания
Напишите так же констрейнт type Number interface, который будет обозначать все численные типы.
*/
