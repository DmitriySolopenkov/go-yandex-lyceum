package main

type World struct {
	Height int // высота сетки
	Width  int // ширина сетки
	Cells  [][]bool
}

// func (w *World) Neighbors(x, y int) int {
// 	count := 0
// 	for i := -1; i <= 1; i++ {
// 		for j := -1; j <= 1; j++ {
// 			// Игнорируем текущую клетку
// 			if i == 0 && j == 0 {
// 				continue
// 			}
// 			// Координаты соседней клетки
// 			newX, newY := x+i, y+j

// 			// Обрабатываем случаи, когда соседняя клетка находится за пределами границы мира.
// 			// Если мир является замкнутой сеткой, обработка за границами необходима.
// 			// В этой реализации предполагается, что мир замкнутый (т.е., клетка справа от крайней правой клетки связана с крайней левой и так далее).
// 			if newX < 0 {
// 				newX = w.Width - 1
// 			} else if newX >= w.Width {
// 				newX = 0
// 			}
// 			if newY < 0 {
// 				newY = w.Height - 1
// 			} else if newY >= w.Height {
// 				newY = 0
// 			}

// 			// Если соседняя клетка жива, увеличиваем счетчик
// 			if w.Cells[newY][newX] {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }

func (w *World) Neighbors(x, y int) int {
	n := 0
	tor := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}
	for _, v := range tor {
		xx, yy := x+v[0], y+v[1]
		if xx < 0 || xx >= w.Height || yy < 0 || yy >= w.Width {
			continue
		}
		if w.Cells[yy][xx] {
			n++
		}
	}

	return n
}

/*
Подсчёт соседей
Напишите недостающий метод для подсчета живых соседних клеток в сетке:
	func (w *World) Neighbors(x, y int) int.
	Всего у клетки может быть максимум восемь соседей.
	Например, клетка в центре (зеленые клетки - живые) имеет три соседа. example image

Примечания
Код программы должен содержать описание струкрутры World:
type World struct { Height int Width int Cells [][]bool }
*/
