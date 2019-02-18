package main

import "fmt"

func main() {

	s := "Sorry Todd, I value my privacy too much."

	oneAtATime(0, s)
}

func oneAtATime(pos int, s string) {
	if pos > len(s) - 1  {
		return
	}
	fmt.Println(string(s[pos]))
	oneAtATime(pos+1, s)

}