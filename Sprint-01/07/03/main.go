/*
Для решения этой задачи с использованием динамического программирования можно использовать подход "набора монет". Мы будем рассматривать каждый размер пиццы как монету определенного номинала, и наша задача состоит в том, чтобы набрать площадь X с минимальной суммой.

Вот решение на языке Go (Golang) для данной задачи:
*/
package main

import (
	"fmt"
)

// func MinPizzaCost(s, m, l, cs, cm, cl, x int) int {
// 	// Создаем срез для хранения минимальных сумм
// 	dp := make([]int, x+1)

// 	// Инициализируем начальные значения для размеров пиццы
// 	dp[0] = 0
// 	for i := 1; i <= x; i++ {
// 		dp[i] = math.MaxInt32
// 	}

// 	// Рассчитываем минимальные суммы для каждой площади
// 	for i := 1; i <= x; i++ {
// 		// Проверяем, можем ли мы купить маленькую пиццу
// 		if i >= s && dp[i-s]+cs < dp[i] {
// 			dp[i] = dp[i-s] + cs
// 		}
// 		// Проверяем, можем ли мы купить среднюю пиццу
// 		if i >= m && dp[i-m]+cm < dp[i] {
// 			dp[i] = dp[i-m] + cm
// 		}
// 		// Проверяем, можем ли мы купить большую пиццу
// 		if i >= l && dp[i-l]+cl < dp[i] {
// 			dp[i] = dp[i-l] + cl
// 		}
// 	}

//		// Возвращаем минимальную сумму для площади X
//		return dp[x]
//	}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func MinPizzaCost(s, m, l, cs, cm, cl, x int) int {
// 	// Создаем двумерный срез dp размером x+1 на 3, где dp[i][j] будет указывать на минимальную сумму денег,
// 	// необходимую для покупки i квадратных сантиметров пиццы типа j (0 - маленькая, 1 - средняя, 2 - большая).
// 	dp := make([][]int, x+1)
// 	for i := 0; i <= x; i++ {
// 		dp[i] = make([]int, 3)
// 	}

// 	// Заполняем таблицу dp построчно
// 	for i := 1; i <= x; i++ {
// 		// Вычисляем стоимость для каждого размера пиццы
// 		costS := cs + dp[i-1][0] // стоимость маленькой пиццы
// 		fmt.Println(costS)
// 		costM := cm + dp[i-1][1] // стоимость средней пиццы
// 		fmt.Println(costM)
// 		costL := cl + dp[i-1][2] // стоимость большой пиццы
// 		fmt.Println(costL)

// 		// Находим минимальную стоимость для текущей клетки двумерного среза
// 		dp[i][0] = min(costS, costM+l) // минимальная стоимость для маленькой пиццы
// 		dp[i][1] = min(costM, costS+l) // минимальная стоимость для средней пиццы
// 		dp[i][2] = min(costL, costM+l) // минимальная стоимость для большой пиццы
// 	}
// 	fmt.Println(dp[x][0])
// 	fmt.Println(dp[x][1])
// 	fmt.Println(dp[x][2])

// 	// Возвращаем минимальную стоимость для покупки x квадратных сантиметров пиццы
// 	return min(dp[x][0], min(dp[x][1], dp[x][2]))
// }

// func MinPizzaCost(s, m, l, cs, cm, cl, x int) int {

// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func minCost(pizza *Pizza, sumArea int) int {
// 	n := len(pizza.cost) // количество пицц
// 	matrix := make([][]int, n+1)
// 	for i := range matrix {
// 		matrix[i] = make([]int, sumArea+1)
// 	}

// 	for item := 1; item <= n; item++ { // переберём все предметы из сундука
// 		for capacity := 1; capacity <= sumArea; capacity++ {
// 			// всё ниже — о рюкзаке вместимостью capacity
// 			maxcostWithoutCurrent := matrix[item-1][capacity] //  максимальная стоимость предыдущих предметов
// 			maxcostWithCurrent := 0                           //  для хранения максимальной стоимости, если положим текущий предмет

// 			weightOfCurrent := pizza.area[item-1] // масса текущего
// 			if capacity >= weightOfCurrent {      // проверяем, влезет ли текущий предмет в рюкзак
// 				// если текущий влез, то смотрим, что ещё взять
// 				maxcostWithCurrent = pizza.cost[item-1]                 // сначала положим текущий предмет
// 				remainingCapacity := capacity - weightOfCurrent         // проверим, осталось ли место
// 				maxcostWithCurrent += matrix[item-1][remainingCapacity] // максимальная стоимость оставшегося места
// 			}

// 			matrix[item][capacity] = max(maxcostWithoutCurrent, maxcostWithCurrent) // выбираем, нужно ли класть текущий
// 		}
// 	}
// 	return matrix[n][sumArea]
// }

// type Pizza struct {
// 	area []int
// 	cost []int
// }

func MinPizzaCost(s int, m int, l int, cs int, cm int, cl int, x int) int {
	table := make([][]int, 3)
	for i := 0; i < 3; i++ {
		table[i] = make([]int, 3)
	}
	table[0][0] = s * cs
	fmt.Println(table)
	table[0][1] = s * cm
	fmt.Println(table)
	table[0][2] = s * cl
	fmt.Println(table)
	table[1][0] = m * cs
	fmt.Println(table)
	table[1][1] = m * cm
	fmt.Println(table)
	table[1][2] = m * cl
	fmt.Println(table)
	table[2][0] = l * cs
	fmt.Println(table)
	table[2][1] = l * cm
	fmt.Println(table)
	table[2][2] = l * cl
	fmt.Println(table)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			min_cost := min(table[i][j], table[i][j], table[i][j])
			table[i][j] = min_cost
		}
	}
	return table[0][0]
	// return 1
}

//	func min(a, b, c int) int {
//		if a <= b && a <= c {
//			return a
//		} else if b <= a && b <= c {
//			return b
//		} else {
//			return c
//		}
//	}
func min(a, b, c int) int {
	minVal := a

	if b < minVal {
		minVal = b
	}

	if c < minVal {
		minVal = c
	}

	return minVal
}

// func MinPizzaCost(s, m, l, cs, cm, cl, x int) int {
// 	table := make([][]int, x+1)
// 	for i := 0; i <= x; i++ {
// 		table[i] = []int{i * s, i * m, i * l}
// 	}
// 	for j := 1; j <= x; j++ {
// 		for k := 1; k < 3; k++ {
// 			if j == 1 || j == 2 {
// 				table[j][k] = min(table[j-1][k], table[j][k-1], table[j+1][k]) + cs
// 			} else if j == 3 {
// 				table[j][k] = min(table[j-1][k], table[j][k-1], table[j+1][k]) + cm
// 			} else if j == 4 {
// 				table[j][k] = min(table[j-1][k], table[j][k-1], table[j+1][k]) + cl
// 			}
// 		}
// 	}
// 	return table[x][0]
// }

func main() {
	s := 4    // площадь маленькой пиццы
	m := 6    // площадь средней пиццы
	l := 12   // площадь большой пиццы
	cs := 40  // стоимость маленькой пиццы
	cm := 60  // стоимость средней пиццы
	cl := 100 // стоимость большой пиццы
	x := 17   // площадь, которую нужно набрать

	minCost := MinPizzaCost(s, m, l, cs, cm, cl, x)
	fmt.Println("Минимальная сумма денег, необходимая для покупки пиццы:", minCost)
	// sumArea := 10 // грузоподъёмность рюкзака
	// pizza := Pizza{
	// 	area: []int{4, 6, 12},
	// 	cost: []int{100, 400, 300, 500}, // стоимость
	// }
	// fmt.Println(minCost(&pizza, sumArea))
}

/*
В этом примере мы передаем параметры размеров пиццы (s, m, l), их стоимости (cs, cm, cl) и площадь, которую нужно набрать (x) в функцию MinPizzaCost. Функция рассчитывает и возвращает минимальную сумму денег для покупки пиццы.

Примечание: В этом решении мы предполагаем, что площади пиццы и необходимая площадь для набора заданы в одинаковых единицах измерения. Если это не так, вам может потребоваться сделать дополнительные преобразования или пересчет единиц.
*/
/*
Golang

Пицца *


В нашем городе есть магазин, в котором продается пицца трех разных размеров:

Маленькая, имеющая площадь - S с м 2 см^2см
2
  и стоимость - CS,
Средняя, имеющая площадь - M с м 2 см^2см
2
  и стоимость - CM и
Большая, имеющая площадь - L с м 2 см^2см
2
  и стоимость - CL
Мы хотим купить X с м 2 см^2см
2
  пиццы.

Какова минимальная сумма денег, необходимая нам для такой покупки?

Напишите функцию MinPizzaCost(s, m, l, cs, cm, cl, x int) int, которая рассчитает и вернёт минимальную сумму.

Используйте динамическое программирование для решения.
*/
