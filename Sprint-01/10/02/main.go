package main

func (s *Stack[T]) Pop() T {
	index := len(s.data) - 1
	val := s.data[index]
	s.data = s.data[:index]
	return val

}

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(val T) {
	s.data = append(s.data, val)
}

/*
Дженерик стэк
Реализуйте дженерик структуру данных стека,
	которая может извлекать и извлекать элементы любого типа.

Функции, которые нужно реализовать Push(val T) Pop() T

Примечания
Структура должна иметь вид Stack[T any] и иметь конструктор NewStack.
*/
