package main

import "fmt"

func main() {
	/*Slices*/

	var slicex []float64
	//slice1[0]=1
	fmt.Println(slicex)
	y := make([]float64, 6)
	fmt.Println(y)
	x := make([]float64, 5, 10)
	fmt.Println(x)
	x[4] = 23
	//fmt.Println(x)
	//x[5]=23
	arr := [6]float64{1, 2, 3, 4, 5, 6}
	w := arr[0:5]
	fmt.Println(w)
	w = x[0:10]
	w[9]=69
	fmt.Println(w)

	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := []int{1,2,3}
	slice4 := make([]int, 2,5)
	//slice4 = slice4[0:5]
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)

}
