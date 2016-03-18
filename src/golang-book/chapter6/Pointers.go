package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}
func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0
	one := func(xPtr *int) {
		*xPtr = 5
	}
	xPtr := new(int)
	*xPtr = 10
	fmt.Println(*xPtr) // x
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}
