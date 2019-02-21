package main

import "fmt"

func main() {

	s := "Sorry Todd, I value my privacy too much.  I will give a brown bag internally."

	oneAtATime(0, s)
}

func oneAtATime(pos int, s string) {
	if pos > len(s) - 1  {
		return
	}
	fmt.Println(string(s[pos]))
	oneAtATime(pos+1, s)

}