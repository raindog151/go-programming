package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {

	inc := 0

	for i := 0; i < 10; i ++ {
		wg.Add(1)
		go func() {
			mu.Lock()
				nv := inc
				nv++
				inc = nv
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(inc)
}
