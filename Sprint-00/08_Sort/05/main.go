func findDisappearedNumbers(nums []int) []int {
	key := make([]int, len(nums)+1)

	// Проходим по входному массиву и увеличиваем соответствующий ключ в key[]
	// для каждого числа, которое встречаем в nums
	for _, n := range nums {
		key[n]++
	}

	res := []int{}
	// Проходим по key[] и добавляем индексы с нулевым значением в res[],
	// таким образом, находим числа, которые отсутствуют в исходном массиве
	for i := 1; i < len(key); i++ {
		if key[i] == 0 {
			res = append(res, i)
		}
	}

	return res
}

func findDisappearedNumbers(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return nil
	}
	res := make([]int, length)

	// Используем res как "отсортированный" массив для записанных значений
	// Проходим по входному массиву и записываем каждое число по его индексу - 1 в res
	for _, v := range nums {
		res[v-1] = v
	}

	// Теперь в res у нас есть массив, в котором пропавшие числа имеют значение 0
	// Проходим по res и "сдвигаем" все значения с нулевым значением в начало массива
	j := 0
	for i := 0; i < length; i++ {
		if res[i] == 0 {
			res[j] = i + 1
			j++
		}
	}

	// Обрезаем res до длины j и возвращаем
	return res[:j]
}