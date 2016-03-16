package main

import "fmt"

const y string = "yyy"

func main() {
	var x string = "Hello World"
	fmt.Println(x)
	f()
	var (
		a string  = "asd"
		b int     = 10
		c float64 = 5.122454
	)
	fmt.Println(a, b, c)
	fmt.Println("asd")
}

func f() {
	fmt.Println(y)
}
