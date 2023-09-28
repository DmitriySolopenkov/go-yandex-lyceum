package main

import "fmt"

func SwapKeysAndValues(m map[string]string) map[string]string {
	res := make(map[string]string)
	for k, v := range m {
		res[v] = k
	}
	return res
}

func main() {
	fmt.Println(SwapKeysAndValues(map[string]string{"Яндекс": "+74957397000", "Музей Яндекса": "+74991101133"}))
}
