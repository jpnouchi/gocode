package main

import "fmt"

func main() {
	fmt.Println("for loop")
	for i := 100; i > 0; i-- {
		if i % 2 == 0 {
			fmt.Println(i, "even")
		}else {
			fmt.Println(i, "odd")
		}
	}

	for i := 1; i < 5; i++ {
		switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		default:
			fmt.Println("zzzzz")
		}
	}

	
}
