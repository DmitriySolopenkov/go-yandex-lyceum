package main

import (
	"errors"
)

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(key T) {
	s.elements = append(s.elements, key)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.elements) == 0 {
		return *new(T), errors.New("Stack is empty")
	}

	index := len(s.elements) - 1
	element := s.elements[index]
	s.elements = s.elements[:index]

	return element, nil
}

// func main() {
// 	stack := Stack[int]{} // Создание стека целых чисел

// 	stack.Push(10)
// 	stack.Push(20)
// 	stack.Push(30)

// 	for {
// 		element, err := stack.Pop()
// 		if err != nil {
// 			break
// 		}
// 		fmt.Println(element)
// 	}
// }

/*
Стек
В программировании стек представляет собой контейнер,
	в который элементы включаются и удаляются только с одного конца,
	называемого вершиной стека.

Когда элемент добавляется в стек,
	он располагается на вершине,
	а при удалении элемента из стека удаляется последний добавленный элемент.

Создайте структуру Stack, которая может содержать элементы любого типа, используя дженерики.

Она должен поддерживать следующие методы:
- Push(key T) - добавление элемента в стек
- Pop() (T, error) - удаление элемента из стека.

В случае успеха должен вернуть значение удаленного элемента, если же список пуст - ошибку.
*/
