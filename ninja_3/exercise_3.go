package main

import (
	"fmt"
	"time"
)

func main() {

	by := 1979

	for by <= int(time.Now().Year()) {
		fmt.Println(by)
		by++
	}
}
