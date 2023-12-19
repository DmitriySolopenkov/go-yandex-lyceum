package main

import "sort"

func SortNames(names []string) {
	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})
}

/*
Сортировка слайса строк
Напишите программу, содержащую функцию `SortNames(names []string)``,
	которая сортирует список имён в алфавитном порядке.
	Если первые символы совпадают, сортировать по вторым, и т.д.

Примечания
Пример отсортированного списка: Аксинья, Арина, Варвара, Есения
*/
