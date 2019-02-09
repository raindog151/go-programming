package main

import "fmt"

func main() {

	a := 4 == 5
	b := 4 <= 5
	c := 4 >= 5
	d := 4 != 5
	e := 4 < 5
	f := 4 > 5

	fmt.Printf("%v %v %v %v %v %v", a, b, c, d, e, f)
}
