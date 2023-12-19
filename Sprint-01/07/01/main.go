package main

import "fmt"

func MaxExpressionValue(nums []int) int {
	n := len(nums)
	maxVal := 0
	// dp := make([][][]int, n)
	// for i := range dp {
	// 	dp[i] = make([][]int, n)
	// 	for j := range dp[i] {
	// 		dp[i][j] = make([]int, n)
	// 	}
	// }

	// first := make([]int, len(nums)+1)
	// for i := len(nums) - 1; i >= 0; i-- {
	// 	first[i] = max(first[i+1], nums[i]) // функция max - возвращает максимальное
	// }

	// second := make([]int, len(nums))
	// for i := len(nums) - 2; i >= 0; i-- {
	// 	second[i] = max(second[i+1], first[i+1]-nums[i])
	// }
	// // здесь будет максимум
	// fmt.Println(second[0])

	for p := n - 1; p >= 0; p-- {
		for q := p + 1; q < n; q++ {
			for r := q + 1; r < n; r++ {
				// dp[p][q][r] = -nums[r] + nums[q] - nums[p]
				for s := r + 1; s < n; s++ {
					maxVal = max(maxVal, nums[s]-nums[r]+nums[q]-nums[p])
				}
				// dp[p][q][r] = maxVal
			}
		}
	}

	// return dp[0][0][0]
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{3, 9, 10, 1, 30, 40}
	fmt.Println(MaxExpressionValue(nums))
	// MaxExpressionValue(nums)
}

/*
Значение выражения


Функция MaxExpressionValue(nums []int) int принимает на вход слайс nums.

Найдите максимальное значение выражения nums[s] — nums[r] + nums[q] — nums[p],
	где p, q, r и s — индексы слайса,
	а s > r > q > p.

Например, для nums := []uint{3, 9, 10, 1, 30, 40} функция должна вернуть значение 46 (поскольку 40 - 1 + 10 - 3 - максимально).

Задачу надо решить, используя принципы динамического программирования.

Примечания
В качестве решения надо отправить функцию MaxExpressionValue и все вспомогательные функции, которые вам потребуются.
*/
