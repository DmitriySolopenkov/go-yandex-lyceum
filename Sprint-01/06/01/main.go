package main

import "errors"

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if nums == nil || n <= 0 {
		return nil, errors.New("invalid input")
	}

	result := make([]int, 0, n)
	for _, num := range nums {
		if num < limit {
			result = append(result, num)
			if len(result) == n {
				break
			}
		}
	}

	return result, nil
}

/*
Лимит по значению


Дан неотсортированный слайс целых чисел.
	Напишите функцию UnderLimit(nums []int, limit int, n int) ([]int, error),
	которая будет возвращать первые n (либо меньше, если остальные не подходят) элементов,
	которые меньше limit. В случае ошибки функция должна вернуть nil и описание ошибки.

Примечания
Функцию main создавать не надо.
*/

/*
	{
		nums:      []int{4, 7, 89, 3, 21, 2, 5, 7, 32, 4, 6, 8, 0, 3, 4, 6, 2, 115, 12},
		n:         5,
		limit:     3,
		expected:  []int{2, 0, 2},
		wantError: false,
	},
	{
		nums:      nil,
		wantError: true,
	},
	{
		nums:      []int{},
		n:         5,
		limit:     3,
		expected:  []int{},
		wantError: false,
	},
	{
		nums:      []int{3, 5, 6},
		n:         5,
		limit:     10,
		expected:  []int{3, 5, 6},
		wantError: false,
	},
	{
		nums:      []int{-13, 0, 6},
		n:         1,
		limit:     -5,
		expected:  []int{-13},
		wantError: false,
	},
	{
		nums:      []int{},
		n:         -1,
		limit:     5,
		wantError: true,
	},
*/
