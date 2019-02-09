package main

import "fmt"

type strelka int

var s strelka = 100
var y int

func main() {

	fmt.Println(s)
	fmt.Printf("%T\n", s)

	s = 42

	fmt.Println(s)

	y = int(s)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
