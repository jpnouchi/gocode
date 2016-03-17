package main

import (
	"fmt"
)

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

func third() {
	fmt.Println("3er")
}

func main() {
	fmt.Println("enter",test(2))

/*	panic("PANIC")
	str := recover() // this will never happen
	fmt.Println(str)*/

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")

}

func test(x int) int {
	defer second()
	first()
	if x == 1 {
		fmt.Println("step 1")
		return 1
	}else if x == 2 {
		fmt.Println("step 2")
		return 2

	}else {
		third()
		return 99
	}
}
