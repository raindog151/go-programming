package main

import "fmt"

type strelka int

var s strelka = 100

func main() {

	fmt.Println(s)
	fmt.Printf("%T\n", s)

	s = 42

	fmt.Println(s)

}
