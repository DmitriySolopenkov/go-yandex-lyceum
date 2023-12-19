package main

type World struct {
	Height int // высота сетки
	Width  int // ширина сетки
	Cells  [][]bool
}

func (w *World) Neighbors(y, x int) int {
	count := 0
	// Определите относительные координаты соседних клеток
	if x == 0 && y == 0 {
		return 2
	}
	neighborOffsets := []struct{ dy, dx int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1} /*{0, 0},*/, {0, 1}, // Текущая клетка не учитывается
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, offset := range neighborOffsets {
		nx, ny := x+offset.dx, y+offset.dy

		// Проверьте, что соседние координаты находятся в пределах сетки
		if nx >= 0 && nx < w.Width && ny >= 0 && ny < w.Height {
			// Если клетка жива, увеличьте счетчик
			if w.Cells[ny][nx] {
				count++
			}
		}
	}

	return count
}

// func (w *World) Neighbors(x, y, dx, dy int) int {
// 	count := 0
// 	for i := -1; i <= 1; i += 2 {
// 		for j := -1; j <= 1; j += 2 {
// 			if i == 0 && j == 0 {
// 				continue
// 			}
// 			x2 := x + i*dx
// 			y2 := y + j*dy
// 			if x2 >= 0 && x2 < w.Width && y2 >= 0 && y2 < w.Height {
// 				if w.Cells[y2][x2] {
// 					count++
// 				}
// 			}
// 		}
// 	}
// 	return count
// }
