package main

import (
	"fmt"
	"time"
)

func TimeAgo(pastTime time.Time) string {
	duration := time.Since(pastTime)

	// Определяем года, месяцы, дни, часы и минуты, прошедшие с момента pastTime
	years := duration.Hours() / (24 * 365)
	months := duration.Hours() / (24 * 30)
	days := duration.Hours() / 24
	hours := duration.Hours()
	minutes := duration.Minutes()

	switch {
	case years >= 1:
		return fmt.Sprintf("%.0f years ago", years)
	case months >= 1:
		return fmt.Sprintf("%.0f months ago", months)
	case days >= 1:
		return fmt.Sprintf("%.0f days ago", days)
	case hours >= 1:
		return fmt.Sprintf("%.0f hours ago", hours)
	default:
		return fmt.Sprintf("%.0f minutes ago", minutes)
	}
	//	if years >= 1 {
	//		return fmt.Sprintf("%.0f years ago", years)
	//	} else if months >= 1 {
	//
	//		return fmt.Sprintf("%.0f months ago", months)
	//	} else if days >= 1 {
	//
	//		return fmt.Sprintf("%.0f days ago", days)
	//	} else if hours >= 1 {
	//
	//		return fmt.Sprintf("%.0f hours ago", hours)
	//	} else {
	//
	//		return fmt.Sprintf("%.0f minutes ago", minutes)
	//	}
}

/*
Сколько времени прошло
Напишите функцию TimeAgo(pastTime time.Time) string, который вычисляет время,
	прошедшее с момента заданного значения time.Time,
	и возвращает удобочитаемую строку, указывающую, как давно это было.

Примечания
Пример работы функции:

    pastTime := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
    result := TimeAgo(pastTime)
    fmt.Println(result) // Output: 1 month ago
Пример работы с часами:

    pastTime := time.Now().Add(-2 * time.Hour)
    timeAgo := TimeAgo(pastTime)
    expected := "2 hours ago"
*/
