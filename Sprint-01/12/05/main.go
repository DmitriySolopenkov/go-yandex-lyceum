package main

import "time"

func NextWorkday(startTime time.Time) time.Time {
	// Define the duration of a day
	oneDay := 24 * time.Hour

	// Get the day of the week for the start time
	weekday := startTime.Weekday()

	// Calculate the number of days until the next workday
	daysUntilNextWorkday := 1
	if weekday == time.Friday {
		// If today is Friday, add 3 days to skip the weekend
		daysUntilNextWorkday = 3
	} else if weekday == time.Saturday {
		// If today is Saturday, add 2 days to skip the weekend
		daysUntilNextWorkday = 2
	}

	// Calculate the next workday by adding the daysUntilNextWorkday to the start time
	nextWorkday := startTime.Add(time.Duration(daysUntilNextWorkday) * oneDay)

	// Check if the next workday is on the weekend, and adjust accordingly
	if nextWorkday.Weekday() == time.Saturday {
		nextWorkday = nextWorkday.Add(oneDay)
	} else if nextWorkday.Weekday() == time.Sunday {
		nextWorkday = nextWorkday.Add(2 * oneDay)
	}

	return nextWorkday
}

/*
Следующий рабочий день
Напишите функцию NextWorkday(start time.Time) time.Time,
	которая вычисляет дату следующего рабочего дня (исключая выходные).

Учитывая дату начала, функция должна возвращать дату следующего рабочего дня.

Примечания
Рабочая неделя начинается с понедельника :)
*/
