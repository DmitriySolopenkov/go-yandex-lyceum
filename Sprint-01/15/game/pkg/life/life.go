package life

import (
	"math/rand"
	"time"
)

type World struct {
	Height int // Высота сетки
	Width  int // Ширина сетки
	Cells  [][]bool
}

// Используйте код из предыдущего урока по игре «Жизнь»
func NewWorld(height, width int) (*World, error) {

	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width) // создаём новый слайс в каждой строке
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}, nil
}

func (w *World) next(x, y int) bool {
	n := w.neighbors(x, y)       // получим количество живых соседей
	alive := w.Cells[y][x]       // текущее состояние клетки
	if n < 4 && n > 1 && alive { // если соседей двое или трое, а клетка жива
		return true // то следующее состояние — жива
	}
	if n == 3 && !alive { // если клетка мертва, но у неё трое соседей
		return true // клетка оживает
	}

	return false // в любых других случаях — клетка мертва
}

func (w *World) neighbors(x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// Игнорируем текущую клетку
			if i == 0 && j == 0 {
				continue
			}
			// Координаты соседней клетки
			newX, newY := x+i, y+j

			// Обрабатываем случаи, когда соседняя клетка находится за пределами границы мира.
			// Если мир является замкнутой сеткой, обработка за границами необходима.
			// В этой реализации предполагается, что мир замкнутый (т.е., клетка справа от крайней правой клетки связана с крайней левой и так далее).
			if newX < 0 {
				newX = w.Width - 1
			} else if newX >= w.Width {
				newX = 0
			}
			if newY < 0 {
				newY = w.Height - 1
			} else if newY >= w.Height {
				newY = 0
			}

			// Если соседняя клетка жива, увеличиваем счетчик
			if w.Cells[newY][newX] {
				count++
			}
		}
	}
	return count
}

func NextState(currentWorld, nextWorld World) {}

// RandInit заполняет поля на указанное число процентов
func (w *World) RandInit(percentage int) {
	// Количество живых клеток
	numAlive := percentage * w.Height * w.Width / 100
	// Заполним живыми первые клетки
	w.fillAlive(numAlive)
	// Получаем рандомные числа
	r := rand.New(rand.NewSource(time.Now().Unix()))

	// Рандомно меняем местами
	for i := 0; i < w.Height*w.Width; i++ {
		randRowLeft := r.Intn(w.Width)
		randColLeft := r.Intn(w.Height)
		randRowRight := r.Intn(w.Width)
		randColRight := r.Intn(w.Height)

		w.Cells[randRowLeft][randColLeft] = w.Cells[randRowRight][randColRight]
	}
}

func (w *World) fillAlive(num int) {
	aliveCount := 0
	for j, row := range w.Cells {
		for k := range row {
			w.Cells[j][k] = true
			aliveCount++
			if aliveCount == num {

				return
			}
		}
	}
}
