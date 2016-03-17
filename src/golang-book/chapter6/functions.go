package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func main() {
	xs := []float64{98,93,77,82,83}
	fmt.Println(average(xs))
	average(xs)
	fmt.Println(f2())
	x,y,z:=f()
	n,err := fmt.Println(x,y,z)
	fmt.Println(n,err)
	val:="value"
	f3(val)
	fmt.Println(val)

}

func f2() (r int) {
	r = 1
	return
}

func f()  int, int ,string  {
	return 5 ,6 ,"9"
}

func f3(aux string) {
	fmt.Println(aux)
	aux="changue"
	fmt.Println(aux)
}