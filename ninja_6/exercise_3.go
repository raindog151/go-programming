package main

import "fmt"

func main() {

	defer foo()
	bar()
}

func foo() {
	fmt.Println("I'm late!")
}

func bar() {
	fmt.Println("I'm early!")
}