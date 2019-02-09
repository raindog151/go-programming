package main

import "fmt"

const (
	a = 1
	b = "Hello"
	c = 5.678

	d int = 2
	e string = "Goodbye"
	f float64 = 11.356
)

func main() {

	fmt.Printf("%T %T %T %T %T %T\n", a, b, c, d, e, f)

}
