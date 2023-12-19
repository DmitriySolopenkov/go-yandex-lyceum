package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Student struct {
	Name          string    `json:"name"`
	AdmissionDate time.Time `json:"admission_date"`
	DaysOnCourse  int       `json:"-"`
}

func processJSON(jsonData []byte) error {
	var students []Student
	date := time.Date(2023, time.October, 2, 0, 0, 0, 0, time.Local)
	// format := "2006-01-02T15:04:05Z"

	// Распаковываем входные данные JSON в структуру
	err := json.Unmarshal(jsonData, &students)
	if err != nil {
		panic(err)
	}

	for i := range students {
		admiss := students[i].AdmissionDate
		diff := date.Sub(admiss)
		students[i].DaysOnCourse = (int(diff.Hours()) / 24)
		fmt.Printf("%v: %v\n", students[i].Name, students[i].DaysOnCourse)
	}

	return nil
}

func main() {
	student := []byte(`[
		{
			"name": "Анна",
			"admission_date": "2023-09-29T00:00:00Z"
		},
		{
			"name": "Иван",
			"admission_date": "2023-09-28T00:00:00Z"
		}
	]`)

	processJSON(student)

}

/*
Мы решили посчитать аналитику и посмотреть, а сколько же с нами учится каждый студент курса по Go - то есть найти кол-во дней, которое он (студент) провел в курсе с момента поступления и до 1 октября 2023 года.

Напишите функцию processJSON(jsonData []byte) error, которая должна принимать данные о студентах в формате JSON, разбирать их и выводить искомое число дней.

Вывод должен быть в формате имяСтудента : количество дней

Формат ввода
[
    {
        "name": "Анна",
        "admission_date": "2023-09-29T00:00:00Z"
    },
    {
        "name": "Иван",
        "admission_date": "2023-09-28T00:00:00Z"
    }
]
Формат вывода
Анна: 2
Иван: 3
Примечания
type Student struct {
    Name         string    `json:"name"`
    AdmissionDate time.Time `json:"admission_date"`
    DaysOnCourse int       `json:"-"`
}

*/
