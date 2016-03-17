package main

import "fmt"

func main() {
	fmt.Println("Enter a Fahrenheit number")
	var inputF float64
	fmt.Scanf("&f",&inputF)

	outpuC := (inputF-32)*5/9
	fmt.Println(outpuC)
}

