package main

type Chest struct {
	val []int
	wt  []int
}

// func Knapsack(chest *Chest, maxWeight int) (int, []int) {
// 	// количество драгоценностей
// 	n := len(chest.cost)

// 	// выделим память под слайсы
// 	matrix := make([][]int, n+1)
// 	for i := range matrix {
// 		matrix[i] = make([]int, maxWeight+1)
// 	}

// 	// переберём все предметы из сундука
// 	for item := 1; item <= n; item++ {
// 		for capacity := 1; capacity <= maxWeight; capacity++ {
// 			// всё ниже — о рюкзаке вместимостью capacity
// 			maxcostWithoutCurrent := matrix[item-1][capacity] //  максимальная стоимость предыдущих предметов
// 			maxcostWithCurrent := 0                        //  для хранения максимальной стоимости, если положим текущий предмет

// 			weightOfCurrent := chest.mass[item-1] // масса текущего
// 			if capacity >= weightOfCurrent {    // проверяем, влезет ли текущий предмет в рюкзак
// 				// если текущий влез, то смотрим, что ещё взять
// 				maxcostWithCurrent = chest.cost[item-1]               // сначала положим текущий предмет
// 				remainingCapacity := capacity — weightOfCurrent     // проверим, осталось ли место
// 				maxcostWithCurrent += matrix[item-1][remainingCapacity] // максимальная стоимость оставшегося места
// 			}

//				matrix[item][capacity] = max(maxcostWithoutCurrent, maxcostWithCurrent) // выбираем, нужно ли класть текущий
//			}
//		}
//		return matrix[n][maxWeight]
//	}
func Knapsack(chest *Chest, maxWeight int) (int, []int) {
	n := len(chest.val)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, maxWeight+1)
	}

	for i := 1; i <= n; i++ {
		for w := 1; w <= maxWeight; w++ {
			if chest.wt[i-1] <= w {
				dp[i][w] = max(chest.val[i-1]+dp[i-1][w-chest.wt[i-1]], dp[i-1][w])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	// Восстановление списка номеров предметов
	taken := make([]int, 0)
	i, w := n, maxWeight
	for i > 0 && w > 0 {
		if dp[i][w] != dp[i-1][w] {
			taken = append(taken, i-1)
			w -= chest.wt[i-1]
		}
		i--
	}

	return dp[n][maxWeight], reverse(taken)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverse(s []int) []int {
	reversed := make([]int, len(s))
	for i, j := 0, len(s)-1; i < len(s); i, j = i+1, j-1 {
		reversed[i] = s[j]
	}
	return reversed
}

/*
Рюкзак *


Решая задачу о рюкзаке, мы вычислили,
	на какую максимальную сумму мы можем унести предметов.

Теперь давайте изменим функцию

Knapsack(chest *Chest, maxWeight int) int

на

Knapsack(chest *Chest, maxWeight int) (int, []int),

чтобы она также возвращала номера предметов,
	которые нужно взять для получения максимальной стоимости.

Примечания
Напомним описание структуры Chest:

type Chest struct {
    val []int
    wt  []int
}
*/
