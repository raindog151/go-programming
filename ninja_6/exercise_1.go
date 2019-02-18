package main

import "fmt"

func main() {

	fmt.Println(foo())
	fmt.Println(bar())

}

func foo() int {
	return 242
}

func bar() (int, string) {

	return 242, "Evil OFF"
}