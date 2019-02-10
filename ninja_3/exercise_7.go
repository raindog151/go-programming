package main

import (
	"fmt"
	"time"
)

func main() {

	y := int(time.Now().Year())

	if y < 2019 {
		fmt.Println("You've magically travelled back in time")
	} else if y > 2019 {
		fmt.Println("You've magically travelled in to the future.")
	} else {
		fmt.Println("Temporal normality achieved.")
	}
}
