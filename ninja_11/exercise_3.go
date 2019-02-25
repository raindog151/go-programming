package main

import "fmt"

type customErr struct {
	e string
}

func (c customErr) Error() string {
	return c.e
}

func main() {
	myErr := customErr{"This is my custom error."}

	foo(myErr)
}

func foo(e error) {
	fmt.Printf("foo() says %T: %v\n", e, e)
}
