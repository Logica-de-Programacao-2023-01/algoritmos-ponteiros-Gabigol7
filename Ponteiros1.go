package main

import "fmt"

func sumOfNaturals(ptr *int, n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	*ptr = sum
}

func main() {
	num := 0
	n := 5
	sumOfNaturals(&num, n)
	fmt.Println(num) // imprime a soma dos 5 primeiros números naturais: 15
}
