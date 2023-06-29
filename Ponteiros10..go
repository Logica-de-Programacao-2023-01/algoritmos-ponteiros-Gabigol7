package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func fillPrimes(slice []int, n int) {
	count := 0
	num := 2
	for count < n {
		if isPrime(num) {
			slice[count] = num
			count++
		}
		num++
	}
}

func main() {
	n := 10
	primes := make([]int, n)
	fillPrimes(primes, n)
	fmt.Println(primes)
}
