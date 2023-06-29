package main

import (
	"fmt"
	"math"
)

func updateWithPiAverage(value *float64) {
	*value = (*value + math.Pi) / 2
}

func main() {
	var num float64 = 3.14
	fmt.Println("Antes da atualização:", num)

	updateWithPiAverage(&num)
	fmt.Println("Após a atualização:", num)
}
