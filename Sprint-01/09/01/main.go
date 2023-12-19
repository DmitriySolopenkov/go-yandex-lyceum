package main

type Book struct {
	Title  string
	Author string
	Year   int
	Genre  string
}

func NewBook(t string, a string, y int, g string) *Book {
	return &Book{Title: t, Author: a, Year: y, Genre: g}
}

/*
Книжная полка
Вы должны создать структуру "Книга" (Book) с следующими полями:

Название (Title) Автор (Author) Год выпуска (Year) Жанр (Genre)
Требуется создать конструктор для структуры "Книга",
	который позволит инициализировать поля структуры при создании экземпляра.
	Конструктор должен принимать значения для всех полей
	и возвращать указатель на созданный экземпляр структуры "Книга".

Примечания
Код программы должен содержать описание струкрутры Book:
type Book struct { Title string Author string Year int Genre string }
*/
