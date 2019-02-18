package main

import "fmt"

func main() {

	s := []int{1,2,3,4,5}

	fmt.Println(foo(s...))
	fmt.Println(bar(s))

}

func foo(i ...int) (int) {

	total := 0

	for _, v := range i {
		total += v
	}

	return total
}

func bar(s []int) (int) {

	total := 0

	for _, v := range s {
		total += v
	}

	return total
}