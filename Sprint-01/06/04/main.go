package main

func Join(nums1, nums2 []int) []int {
	result := make([]int, len(nums1)+len(nums2))
	copy(result, nums1)
	copy(result[len(nums1):], nums2)
	return result
}

/*
Слияние

Даны 2 слайса целых чисел nums1 и nums2.
	Создайте функцию Join(nums1, nums2 []int) []int,
	которя создаст новый слайс емкостью,
	вмещающей в себя ровно два слайса (ёмкость должна быть равна его длине).
	Скопируйте в него сначала значения nums1 затем nums2 и верните его.

Примечания
Функцию main создавать не надо.
*/
