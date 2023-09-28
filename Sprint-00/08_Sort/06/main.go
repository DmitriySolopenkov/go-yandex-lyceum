func singleNumber(nums []int) int {
	sort.Ints(nums) // Сортируем массив

	sum := 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 { // Если индекс четный, добавляем число к сумме
			sum += nums[i]
		} else { // Если индекс нечетный, вычитаем число из суммы и проверяем на равенство нулю
			sum -= nums[i]
			if sum != 0 {
				return nums[i-1] // Возвращаем число, предшествующее текущему числу
			}
		}
	}

	return sum // Возвращаем результат, который будет равен одиночному числу
}

func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i] // Используем операцию XOR для обнаружения одиночного числа
	}
	return res
}
