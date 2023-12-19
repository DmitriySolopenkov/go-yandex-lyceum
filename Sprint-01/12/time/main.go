package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Current time:", now)

	t1 := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Println("Difference between t1 and t2:", diff)

	// now := time.Now()

	// Вот как может быть отформатировано время
	fmt.Println("1. Время в формате RFC3339:", now.Format(time.RFC3339))
	fmt.Println("2. Полная дата и время:", now.Format("2006-01-02 15:04:05"))
	fmt.Println("3. Краткая дата:", now.Format("2006-01-02"))
	fmt.Println("4. Время в 24-часовом формате:", now.Format("15:04:05"))
	fmt.Println("5. Время в 12-часовом формате с AM/PM:", now.Format("03:04 PM"))
	fmt.Println("6. День недели:", now.Format("Monday"))
	fmt.Println("7. Сокращённый месяц:", now.Format("Jan"))

	fmt.Println("Start")
	time.Sleep(5 * time.Second)
	fmt.Println("End")

	_ = time.NewTimer(5 * time.Second) // создается таймер  на 5 секунд
	duration := 5 * time.Second
	fmt.Println("Duration:", duration)

	// Выполняем действия в течение заданного промежутка времени
	fmt.Println("Start")
	time.Sleep(duration)
	fmt.Println("End")

	// Вычисляем разницу между моментами времени
	t3 := time.Now()
	time.Sleep(duration)
	t4 := time.Now()
	diff = t4.Sub(t3)
	fmt.Println("Difference:", diff)

	ticker := time.Tick(1 * time.Second)
	for t := range ticker {
		fmt.Println("Tick at", t)
	}

}
