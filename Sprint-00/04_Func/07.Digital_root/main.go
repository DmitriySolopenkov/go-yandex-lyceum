package main

import "fmt"

func CalculateDigitalRoot(n int) int {
	var d []int
	var res int
	for n != 0 {
		d = append(d, n%10)
		n /= 10
	}
	for _, i := range d {
		res += i
	}
	if res/10 != 0 {
		return CalculateDigitalRoot(res)
	}
	return res
}

func main() {
	fmt.Println(CalculateDigitalRoot(12345)) // Например, цифровой корень числа 12345 равен 1 + 2 + 3 + 4 + 5 = 15, а затем 1 + 5 = 6
}
