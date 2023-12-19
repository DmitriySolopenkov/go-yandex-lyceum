package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type World struct {
	Height int // высота сетки
	Width  int // ширина сетки
	Cells  [][]bool
}

func NewWorld(height, width int) *World {
	// создаём тип World с количеством слайсов hight (количество строк)
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width) // создаём новый слайс в каждой строке
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}
}

func (w *World) Neighbours(x, y int) int {
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

func (w *World) Next(x, y int) bool {
	n := w.Neighbours(x, y)      // получим количество живых соседей
	alive := w.Cells[y][x]       // текущее состояние клетки
	if n < 4 && n > 1 && alive { // если соседей двое или трое, а клетка жива
		return true // то следующее состояние — жива
	}
	if n == 3 && !alive { // если клетка мертва, но у неё трое соседей
		return true // клетка оживает
	}

	return false // в любых других случаях — клетка мертва
}

func NextState(oldWorld, newWorld *World) {
	// переберём все клетки, чтобы понять, в каком они состоянии
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			// для каждой клетки получим новое состояние
			newWorld.Cells[i][j] = oldWorld.Next(j, i)
		}
	}
}

// Начальное состояние
func (w *World) Seed() {
	// снова переберём все клетки
	for _, row := range w.Cells {
		for i := range row {
			//rand.Intn(10) возвращает случайное число из диапазона	от 0 до 9
			if rand.Intn(10) == 1 {
				row[i] = true
			}
		}
	}
}

func (w *World) SaveState(filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for item, row := range w.Cells {
		for _, cell := range row {
			var state byte
			if cell {
				state = '1'
			} else {
				state = '0'
			}
			_, err := file.Write([]byte{state})
			if err != nil {
				return err
			}
		}
		if item < len(w.Cells)-1 {
			_, err := file.Write([]byte{10})
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func (w *World) LoadState(filename string) error {
	// Открываем файл для чтения
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var cells [][]bool
	var width int

	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if lineNum == 0 {
			width = len(line)
		} else if len(line) != width {
			return errors.New("Inconsistent line length in the file")
		}

		var row []bool
		for _, char := range line {
			if char == '1' {
				row = append(row, true)
			} else if char == '0' {
				row = append(row, false)
			} else {
				return errors.New("Invalid character in the file")
			}
		}
		cells = append(cells, row)
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	w.Height = len(cells)
	w.Width = width
	w.Cells = cells

	return nil
}

func main() {
	// зададим размеры сетки
	height := 10
	width := 10
	// объект для хранения текущего состояния сетки
	currentWorld := NewWorld(height, width)
	// объект для хранения следующего состояния сетки
	nextWorld := NewWorld(height, width)
	// установим начальное состояние
	currentWorld.Seed()
	for { // цикл для вывода каждого состояния
		// выведем текущее состояние на экран
		fmt.Println(currentWorld)
		// рассчитываем следующее состояние
		NextState(currentWorld, nextWorld)
		// изменяем текущее состояние
		currentWorld = nextWorld
		// делаем паузу
		time.Sleep(100 * time.Millisecond)
		// специальная последовательность для очистки экрана после каждого шага
		fmt.Print("\033[H\033[2J")
	}
}

/*
Сохранение состояния
Напишите новый метод func (w *World) SaveState(filename string) error
	для сохранения текущего состояния сетки в файл.
	Метод должен создавать новый файл и записать данные в бинарном виде, например:
110011
100101

Примечания
Код программы должен содержать описание струкрутры World:
type World struct { Height int Width int Cells [][]bool }
*/
