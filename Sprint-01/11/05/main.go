package main

import (
	"errors"
	"sort"
	"strconv"
)

type CompanyInterface interface {
	AddWokerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

type Worker struct {
	Name       string
	Position   string
	Salary     uint
	Experience uint
}

type Company struct {
	Workers []*Worker
}

func NewCompany() *Company {
	return &Company{Workers: make([]*Worker, 0)}
}

func (c *Company) AddWokerInfo(name, position string, salary, experience uint) error {
	allowedPositions := map[string]bool{
		"директор":       true,
		"зам. директора": true,
		"начальник цеха": true,
		"мастер":         true,
		"рабочий":        true,
	}

	if !allowedPositions[position] {
		return errors.New("недопустимая должность")
	}

	worker := &Worker{
		Name:       name,
		Position:   position,
		Salary:     salary,
		Experience: experience,
	}

	c.Workers = append(c.Workers, worker)
	return nil
}

func (c *Company) SortWorkers() ([]string, error) {
	sort.SliceStable(c.Workers, func(i, j int) bool {
		if c.Workers[i].Salary*c.Workers[i].Experience != c.Workers[j].Salary*c.Workers[j].Experience {
			return c.Workers[i].Salary*c.Workers[i].Experience > c.Workers[j].Salary*c.Workers[j].Experience
		}

		if c.Workers[i].Salary*c.Workers[i].Experience == c.Workers[j].Salary*c.Workers[j].Experience {
			return c.Workers[i].Salary*c.Workers[i].Experience < c.Workers[j].Salary*c.Workers[j].Experience
		}

		if c.Workers[i].Position != c.Workers[j].Position {
			return c.Workers[i].Position > c.Workers[j].Position
		}
		return c.Workers[i].Name < c.Workers[j].Name
	})

	result := make([]string, len(c.Workers))
	for i, worker := range c.Workers {
		result[i] = worker.Name + " - " + strconv.FormatUint(uint64(worker.Salary*worker.Experience), 10) + " - " + worker.Position
	}

	return result, nil
}

/*
Сортировка сотрудников предприятия

На предприятии работают несколько cсотрудников.
Каждый из них имеет свою должность,
	фиксированную месячную заработную плату и стаж работы.

Напишите программу в котором тип Company реализует следующий интерфейс:

type CompanyInterface interface {
	AddWokerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

AddWokerInfo - метод добавления информации о новых сотрудниках,
	где name - имя сотрудника,
	position - должность,
	salary - месячная зароботная плата,
	experience - стаж работы (месяцев).

SortWorkers - метод сортировки, который сортирует список сотрудников по следующим свойствам:
	доход за время работы на предприятии(по убыванию),
	должность (от высокой)
	и возвращает слайсл формата: имя - доход - должность.


Допустимые должности в порядке убывания: "директор", "зам. директора", "начальник цеха", "мастер", "рабочий".

Примечания
Пример отсортированного вывода:
Михаил - 12000 - директор
Андрей - 11800 - мастер
Игорь - 11000 - зам. директора


--- FAIL: TestWorkerSort (0.00s)
    source_test.go:61: SortWorkers failed. Expected:
	[Михаил - 12000 - директор Андрей - 10800 - мастер Игорь - 6480 - зам. директора Николай - 2880 - начальник цеха Виктор - 2880 - рабочий], Got:
	[Виктор - 2880 - рабочий Николай - 2880 - начальник цеха Игорь - 6480 - зам. директора Андрей - 10800 - мастер Михаил - 12000 - директор]
    source_test.go:61: SortWorkers failed. Expected:
	[Максим - 7680 - рабочий Иван - 2880 - начальник цеха Андрей - 2400 - директор], Got:
	[Андрей - 2400 - директор Иван - 2880 - начальник цеха Максим - 7680 - рабочий]
FAIL
*/
