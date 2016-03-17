package main

import "fmt"

func makeEvenGenerator() func() (uint,string) {
	i := uint(0)
	return func() (ret uint,asd string) {
		ret = i
		asd = "asd"
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1,1))

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4
	fmt.Println(factorial(3))

}
