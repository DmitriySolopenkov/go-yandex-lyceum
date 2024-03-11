package once

import (
	"sync"
)

var (
	once     sync.Once
	counter  int
	IncrFunc func()
)

// initialize увеличивает счётчик только один раз
func initialize() {
	counter++
}

// Increment вызывает initialize только один раз
func Increment() {
	once.Do(initialize)
}

// GetCounter возвращает текущее значение счётчика
func GetCounter() int {
	return counter
}
