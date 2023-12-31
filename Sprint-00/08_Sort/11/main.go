func twoSum(nums []int, target int) []int {
	// Инициализируем два указателя, которые указывают на начало и конец массива
	left, right := 0, len(nums)-1

	// Итерируемся до тех пор, пока указатели не пересекутся или не найдём пару
	for left < right {
		// Вычисляем сумму элементов, которые находятся под указателями left и right
		sum := nums[left] + nums[right]
		// Если сумма равна целевому значению, мы нашли пару и возвращаем её
		if sum == target {
			return []int{nums[left], nums[right]}
			// Если сумма меньше целевого значения, увеличиваем left, чтобы увеличить сумму
		} else if sum < target {
			left++
			// Если сумма больше целевого значения, уменьшаем right, чтобы уменьшить сумму
		} else {
			right--
		}
	}

	// Если пара не найдена, возвращаем nil
	return nil
}
