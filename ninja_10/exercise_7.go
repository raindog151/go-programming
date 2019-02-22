package main

import "fmt"

func main() {

	ch := make(chan int)

	for i := 0; i < 10; i++ {
		go populate(ch)
	}

	// the following results in deadlocks and i don't know why

	// for v := range ch {

	for j := 0; j < 100; j++ {
		fmt.Println(<-ch)
	}
}

func populate(c chan int) {

	for i := 0; i < 10; i++ {
		c <- i
	}

}