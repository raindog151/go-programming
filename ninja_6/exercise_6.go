package main

import "fmt"

func main() {

	func() {
		s := []int{1,2,3,4,5}
		for _, v := range s {
			fmt.Println(v)
		}
	}()

}
