package main

// сундук с драгоценностями
type Chest struct {
	cost []int // стоимость предметов
	mass  []int // масса предметов
}

// w := 10       // грузоподъёмность рюкзака
// ch := Chest{
//     cost: []int{100, 400, 300, 500}, // стоимость
//     mass:  []int{5, 4, 6, 3},         // масса
// }

func Knapsack(chest *Chest, maxWeight int) int {
	// количество драгоценностей
	n := len(chest.cost)

	// выделим память под слайсы
	matrix := make([][]int, n+1)
	for i := range matrix {
		matrix[i] = make([]int, maxWeight+1)
	}

	// переберём все предметы из сундука
	for item := 1; item <= n; item++ { 
		for capacity := 1; capacity <= maxWeight; capacity++ {
			// всё ниже — о рюкзаке вместимостью capacity
			maxcostWithoutCurrent := matrix[item-1][capacity] //  максимальная стоимость предыдущих предметов
			maxcostWithCurrent := 0                        //  для хранения максимальной стоимости, если положим текущий предмет

			weightOfCurrent := chest.mass[item-1] // масса текущего
			if capacity >= weightOfCurrent {    // проверяем, влезет ли текущий предмет в рюкзак
				// если текущий влез, то смотрим, что ещё взять
				maxcostWithCurrent = chest.cost[item-1]               // сначала положим текущий предмет
				remainingCapacity := capacity — weightOfCurrent     // проверим, осталось ли место
				maxcostWithCurrent += matrix[item-1][remainingCapacity] // максимальная стоимость оставшегося места
			}

			matrix[item][capacity] = max(maxcostWithoutCurrent, maxcostWithCurrent) // выбираем, нужно ли класть текущий
		}
	}
	return matrix[n][maxWeight]
}
