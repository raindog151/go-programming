package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

		q <- 1
		close(c)
	}()

	return c
}

func receive(c, q <-chan int) {
	// we need to have this infinite loop here to 'wait' for the gen func to populate the
	// channel which is weird for my brain
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q: //
			return
		}
	}
}