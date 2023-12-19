package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)

	// Создаем срез для хранения фамилий
	names := make([]string, n)

	// Считываем фамилии
	in := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {

		line, _ := in.ReadString('\n')
		names[i] = line
	}
	// Сортируем фамилии в лексикографическом порядке
	sort.Strings(names)

	/////
	prefixes := make([]string, 0)
	for {
		prefix, _ := in.ReadString('\n')
		prefix = strings.TrimSpace(prefix)

		if strings.TrimSpace(prefix) == "" {
			break
		}
		prefixes = append(prefixes, prefix)
	}
	var res []string
	for _, prefix := range prefixes {
		found := false
		for i, person := range names {
			if strings.HasPrefix(person, prefix) {
				res = append(res, names[i])
				found = true
				break
			}

		}
		if !found {
			res = append(res, "Не найдено")
		}

	}
	for _, s := range res {
		fmt.Print(s)
	}
}
