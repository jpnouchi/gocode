package main

import "fmt"

func main() {
	fmt.Println("maps")
	var x map[string]int
	x=make(map[string]int)
	x["key"] = 10
	fmt.Println(x)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])
	delete(elements, "Li")
	fmt.Println(elements)
	fmt.Println(elements["Li"])
	name, ok := elements["Ne"]
	fmt.Println(name, ok)

	z := [6]string{"a","b","c","d","e","f"}
	fmt.Println(z[2:4])

	w := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}

	max:=0
	for _,value := range w  {
		if value > max {
			max=value
		}
	}
	fmt.Println("max values is",max)

	array := make([]int, 3, 9)
	fmt.Println("size slice 3 9",len(array))

}

