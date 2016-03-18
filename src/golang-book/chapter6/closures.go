package main

import "fmt"

func makeEvenGenerator() func() (uint, string) {
	i := uint(0)
	return func() (ret uint, asd string) {
		ret = i
		asd = "asd"
		i += 2
		return
	}
}

func makeOddGenerator() func() (uint) {
	k := uint(1)
	return func() (uint) {
		ret := k
		k += 2
		return ret
	}
}

func progress2() (ret int) {
	z := 0
	ret = z + 2
	return
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x - 1)
}

func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))

	nextEven := makeEvenGenerator()
	fmt.Println("par")
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4
	fmt.Println("factorial")
	fmt.Println(factorial(3))
	fmt.Println("impar")
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println("test")
	test := progress2()
	fmt.Println(test)
	fmt.Println(test)
	fmt.Println(test)

	fmt.Println("fibo")
	fmt.Println(3)

}

func fibonacci(val int) int {
	switch val {
	case 0: return 0
	case 1: return 1
	default:
		return fibonacci(val-1) + fibonacci(val-2)
	}
}
