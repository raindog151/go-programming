package main

import "fmt"

type person struct {
	first string
	last string
	flavors []string
}

func main() {

	p1 := person {
		first: "Hal",
		last: "Incandenza",
		flavors: []string{"vanilla", "strawberry"},
	}

	p2 := person {
		first: "Joelle",
		last: "Van Dyne",
		flavors: []string{"pistachio"},
	}

	fmt.Printf("%s %s - ", p1.first, p1.last)
	for _, v := range p1.flavors {
		fmt.Printf("%s ", v)
	}
	fmt.Printf("\n")

	fmt.Printf("%s %s - ", p2.first, p2.last)
	for _, v := range p2.flavors {
		fmt.Printf("%s ", v)
	}
	fmt.Printf("\n")

}
