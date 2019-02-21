package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	inc := 0

	for i := 0; i < 10; i ++ {
		wg.Add(1)
		go func() {
			nv := inc
			runtime.Gosched()
			nv++
			inc = nv
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(inc)
}
