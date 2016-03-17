package main

import (
	"fmt"
)

func main() {
	var x [5]float64
	x[4] = 100
	fmt.Println(x)
	numbers := [5]float64{
		45,
		25,
		65,
		8,
		55,
	}
	fmt.Println(numbers)

	var total float64
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	fmt.Println("total", total)

	total = 0
	for z , value := range numbers{
		total += value
		fmt.Println(z)
	}
	fmt.Println("total", total)

	total = 0
	for _, value := range numbers{
		total += value
	}
	fmt.Println("total", total)

}


