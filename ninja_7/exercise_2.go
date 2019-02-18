package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func main() {

	me := person{
		first: "Dan",
		last: "Nagy",
		age: 39,
	}

	fmt.Println(me)
	changeMe(&me)
	fmt.Println(me)

}

func changeMe(p *person) {
	p.age = 40
}