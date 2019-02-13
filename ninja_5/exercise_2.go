package main

import "fmt"

type person struct {
	first string
	last string
	flavors []string
}

func main() {

	var people = make(map[string]person)

	people["Incandenza"] = person {
		first: "Hal",
		last: "Incandenza",
		flavors: []string{"vanilla", "strawberry"},
	}

	people["Van Dyne"] = person {
		first: "Joelle",
		last: "Van Dyne",
		flavors: []string{"pistachio"},
	}

	for _, v := range people {
		fmt.Printf("%s %s - ", v.first, v.last)
		for _, v := range v.flavors {
			fmt.Printf("%s ", v)
		}
		fmt.Printf("\n")
	}
}
