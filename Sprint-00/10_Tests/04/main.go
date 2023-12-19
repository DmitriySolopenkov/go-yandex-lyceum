package main

import (
	"errors"
	"fmt"
)

const (
	place  = 0
	target = 3
)

type Grasshopper struct { // знает своё местоположение на линейке
	position int  // Текущее положение кузнечика на линейке
	target   int  // Положение зернышка на линейке
	eaten    bool // Переменная, чтобы отслеживать, съел ли кузнечик зернышко
}

func (g *Grasshopper) WhereAmI() int {
	return g.position
}

func (g *Grasshopper) Jump() (int, error) {
	// Проверяем, съел ли кузнечик зернышко
	if g.position == g.target {
		g.eaten = true
	}
	if g.eaten {
		return -1, errors.New("Кузнечик уже съел зернышко")
	}

	// Вычисляем расстояние до зернышка и ограничиваем прыжок в 5 см
	jumpDistance := g.target - g.position

	switch {
	case jumpDistance > 5:
		jumpDistance = 5
	case jumpDistance < -5:
		jumpDistance = -5
	}

	g.position += jumpDistance

	// Проверяем, съел ли кузнечик зернышко
	if g.position == g.target {
		g.eaten = true
	}
	return g.position, nil
}

type Jumper interface {
	WhereAmI() int      // выводит текущее положение кузнечика на линейке
	Jump() (int, error) // кузнечик прыгает к зерну. Выводит новое положение кузнечика, или ошибку, если он уже ест зерно
}

func PlaceJumper(place, target int) Jumper {
	return &Grasshopper{
		position: place,
		target:   target,
		eaten:    false,
	}
}

func main() {
	g := PlaceJumper(place, target)
	fmt.Println(g.WhereAmI())
	for {
		currPlace, err := g.Jump()
		if err != nil {
			break
		}
		fmt.Println(currPlace, err)
	}
}
