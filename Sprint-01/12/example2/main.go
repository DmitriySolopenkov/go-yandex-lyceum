package main

import (
	"fmt"
	"time"
)

func main() {
	// Устанавливаем интервал времени для таймера (в данном случае, 5 секунд)
	interval := 5 * time.Second

	// Создаём новый таймер с указанным интервалом
	timer := time.NewTimer(interval)

	fmt.Println("Задача будет выполнена через", interval)

	// Ожидаем события от таймера (пока не прошло 5 секунд)
	<-timer.C

	fmt.Println("Задача выполнена!")
}
