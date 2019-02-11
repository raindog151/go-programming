package main

import "fmt"

func main() {

	p1 := []string{"Jamees", "Bond", "Shaken not stirred"}
	p2 := []string{"Miss", "Moneypenny", "Hello James"}

	arr := [][]string{p1, p2}

	for i, v := range arr {
		fmt.Println(i, v)
	}
}
