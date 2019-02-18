package main

import "fmt"

func main() {

	ladybird := boof()

	fmt.Println(ladybird())
	fmt.Println(ladybird())
	fmt.Println(ladybird())
	fmt.Println(ladybird())

	liz := boof()
	fmt.Println(liz())
	fmt.Println(liz())
}

func boof() func() (int) {

	b := 0
	return func() int {
		b++
		return b
	}

}