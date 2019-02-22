package main

import "fmt"

func main() {

	ch := make(chan int)

	go populate(ch)

	depopulate(ch)

}

func populate(c chan int) {

	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func depopulate(c chan int) {

	for v := range c {
		fmt.Println(v)
	}

}