package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func main() {

	p := person {
		first: "Dan",
		last: "Nagy",
		age: 39,
	}

	p.speak()
}

func (p person) speak() {
	fmt.Printf("Hello my name is %s %s, I am %d years old.\n", p.first, p.last, p.age)
}