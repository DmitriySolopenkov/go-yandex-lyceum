package main

func Mix(nums []int) []int {
	n := len(nums) / 2
	result := make([]int, 0, len(nums))
	for i := 0; i < n; i++ {
		result = append(result, nums[i], nums[i+n])
	}
	return result
}

/*
Слияние двух частей

Дан слайс nums,
	состоящий из 2n элементов в формате [x0,x1,...,xn,y0,y1,...,yn].
	Создайте функцию Mix(nums []int) []int,
	которая вернёт слайс, содержащий значения в следующем порядке: [x0,y0,x1,y1,...,xn,yn].

Примечания
Функцию main создавать не надо.
*/
