package main

import "fmt"

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string

}
func (p *Person)talk()  {
	fmt.Println("My name is",p.Name)
}

func main() {
	fmt.Println("embedded types")

	p:=&Person{"jpn"}
	p.talk()
	fmt.Println(p)

	a:=new(Android)
	a.Name="luis"
	a.talk()

}
