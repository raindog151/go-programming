package main

import "fmt"

func main() {

	favSport := "Hockey"

	switch favSport {
	case "Football" :
		fmt.Println("I like the Cleveland Browns!")
	case "Baseball":
		fmt.Println("Baseball?  Can't be bothered.")
	case "Hockey":
		fmt.Println("Let's go Wild!")
	default:
		fmt.Printf("I don't know how to play %s!\n", favSport)
	}
}
