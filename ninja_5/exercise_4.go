package main

import "fmt"

func main() {

	car := struct {
		make string
		model string
		year int
		color string
	}{
		make: "Honda",
		model: "Element",
		year: 2004,
		color: "Green",
	}

	fmt.Println(car)
}
