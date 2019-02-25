package main

import "fmt"

func main() {

	fmt.Println(Foo(4, 5))

}

func Foo(i, j int) int {
	return (i + j)
}
