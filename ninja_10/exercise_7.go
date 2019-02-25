package main

import "fmt"

func main() {

	ch := make(chan int)

	for i := 0; i < 10; i++ {
		go populate(ch)
	}

	// the following results in deadlocks and i don't know why
	// update: the reason why is since we don't close the channel (would need to use waitgroups or something
	// the range function sits and waits for input indefinitely and then poops the bed

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