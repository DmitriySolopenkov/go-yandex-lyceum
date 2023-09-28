func moveZeroes(nums []int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}

	// Заполняем оставшиеся элементы массива нулями
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}

	return nums
}

func moveZeroes(nums []int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j { // Проверяем, что i и j не равны
				nums[j] = nums[i]
				nums[i] = 0
			}
			j++
		}
	}

	return nums
}
