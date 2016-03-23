package main

import ("fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Join([]string{"A","B"},"-"))
	arr:=[]byte("test")
	str:=string([]byte{'t','e','s','t'})
	fmt.Println(arr)
	fmt.Println(str)
}

