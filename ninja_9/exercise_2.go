package main

import "fmt"

type Person struct {
	First string
	Last string
	Age int
}

type Human interface {
	speak(s string)
}

func (p *Person) speak(s string) {
	fmt.Printf("%s %s says %s", p.First, p.Last, s)
}

func saySomething(h Human, s string) {
	h.speak(s)
}

func main() {

	me := Person{
		First: "Dan",
		Last: "Nagy",
		Age: 40,
	}

	saySomething(&me, "Hello")

	// I Fail here:
	saySomething(me, "Hello")

}
