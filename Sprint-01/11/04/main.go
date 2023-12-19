package main

import "sort"

func SortByFreq(str string) string {
	charCount := make(map[rune]int)

	for _, char := range str {
		charCount[char]++
	}

	charList := make([]struct {
		char  rune
		count int
	}, 0)

	for char, count := range charCount {
		charList = append(charList, struct {
			char  rune
			count int
		}{char, count})
	}

	sort.Slice(charList, func(i, j int) bool {
		if charList[i].count == charList[j].count {
			return charList[i].char < charList[j].char
		}
		return charList[i].count < charList[j].count
	})

	result := make([]rune, 0)
	for _, item := range charList {
		for i := 0; i < item.count; i++ {
			result = append(result, item.char)
		}
	}

	return string(result)

}

/*
Сортировка символов по частоте

Дана строка с символами из набора алфавита.

Напишите программу с функцией SortByFreq(str string) string,
	которая будет сортировать символы из строки по возрастанию с учетом частоты повторения.

Символы с наименьшим количеством вхождений должны идти в начале,
	а символы с наибольшей частотой - в конце.

В случае одинаковой частоты символов, они должны быть отсортированы в алфавитном порядке.

Примечания
Пример:
Вход: "abbbzzzat"
Выход: "taabbbzzz"
*/
