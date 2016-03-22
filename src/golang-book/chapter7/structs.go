package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func (c *Circle)area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	fmt.Println("structs")
	var c Circle
	fmt.Println(c)
	c.r = 5
	fmt.Println(c)

	ci := new(Circle)
	ci.x = 1
	ci.r = 1
	fmt.Println(ci, "area is", ci.area())

	cir := Circle{r:25, x:1, y:1}
	cir.x = 2
	fmt.Println(cir)

	circ := Circle{2, 2, 125}
	circ.x = 3
	fmt.Println(circ)

	circle := &Circle{0, 0, 15}
	circle.x = 4
	fmt.Println(circle)

	area := func(c Circle) float64 {
		return math.Pi * c.r * c.r
	}
	fmt.Println("area circle", area(*circle))
	fmt.Println("area circ", area(circ))
	fmt.Println("area ci", area(*ci))
	fmt.Println("area c", area(c))

	swap := func(c *Circle) {
		fmt.Println("after", c)
		c.r = 3
		fmt.Println("before")
	}
	swap(circle)
	fmt.Println("result circle", circle)
	swap(&circ)
	fmt.Println("result circ", circ)
	swap(ci)
	fmt.Println("result circ", ci)
	swap(&c)
	fmt.Println("result c", c)

}
