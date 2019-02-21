package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wc := 20

	wg.Add(wc)

	fmt.Println(wg)

	for i := 0; i < 10; i++ {
		go a()
		go b()
	}

	wg.Wait()

	fmt.Println(wg)

}

func a() {
	fmt.Println(("I'm func a."))
	wg.Done()
}

func b() {
	fmt.Println(("I'm func b."))
	wg.Done()
}