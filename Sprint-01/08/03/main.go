package main

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func (w *World) String() string {
	var output string
	for _, row := range w.Cells {
		for _, cell := range row {
			if cell {
				output += "▓" // Живая клетка
			} else {
				output += " " // Мертвая клетка
			}
		}
		output += "\n"
	}
	return output
}

// func main() {
// 	// Пример использования метода String()
// 	height := 5
// 	width := 5
// 	cells := [][]bool{
// 		{true, false, true, false, true},
// 		{false, true, false, true, false},
// 		{true, false, true, false, true},
// 		{false, true, false, true, false},
// 		{true, false, true, false, true},
// 	}

// 	world := World{
// 		Height: height,
// 		Width:  width,
// 		Cells:  cells,
// 	}

// 	fmt.Println(world)
// }

/*
Вывод на экран
В приведенном коде игры текущее состояние сетки выводится на экран следующим образом:
	fmt.Println(currentWorld), что не совсем наглядно.
	Чтобы получить наглядное изображение нашей сетки мы можем отрисовывать клетки разного цвета,
	например:
brownSquare := "\xF0\x9F\x9F\xAB"
greenSquare := "\xF0\x9F\x9F\xA9"
Для этого, созданный нами тип type World struct {
Height int // высота сетки
Width int // ширина сетки
Cells [][]bool
}
должен иметь метод String(),
	который будет формировать строку для отображения.
	Напишите данный метод и попробуйте запустить всю программу, поэкспериментируйте с разными символами и цветами.

Примечания
Код программы должен содержать описание струкрутры World:
type World struct { Height int Width int Cells [][]bool }
*/
