package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {

	var inc int64 = 0

	for i := 0; i < 10; i ++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&inc, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(inc)
}
