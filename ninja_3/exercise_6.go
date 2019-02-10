package main

import (
	"fmt"
	"time"
)

func main() {

	y := int(time.Now().Year())

	if y == 2019 {
		fmt.Println("Temporal normality achieved.")
	}
}
