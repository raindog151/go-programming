package main

import "fmt"

func main() {

	x := boofer(woofer)
	fmt.Println(x)
}

func boofer(f func() string) string {
	return fmt.Sprintf("%v Boof!", f())
}

func woofer() string {
	return fmt.Sprintf("My dog Ladybird says: Woof!")
}