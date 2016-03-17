package chapter3

import "fmt"

func main() {
	fmt.Println("enter a number")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}
