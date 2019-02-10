package main

import (
	"fmt"
	"time"
)

func main() {

	by := 1979

	for {
		if by <= int(time.Now().Year()) {
			fmt.Println(by)
			by++
		} else {
			break
		}
	}
}

