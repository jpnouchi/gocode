package chapter3

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
	var z = 4
	var w int = 5
	y := "jkl"
	r := 1010
	r += 1
 	fmt.Println(z)
	fmt.Println(w)
	fmt.Println(y)
	fmt.Println(r)


}

func f() {
	fmt.Println(y)
}
