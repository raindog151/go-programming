package main

import "fmt"

func main() {

	f := bark()

	fmt.Printf("%v", f())
}

func bark() func() string {
	return func() string {
		return fmt.Sprintln("Woof!")
	}
}