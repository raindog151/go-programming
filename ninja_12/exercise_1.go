package dog

import "fmt"

func main() {

	age := 2

	fmt.Printf("Ladybird is %d years old or %d in dog years.", age, Years(age))
}

// Take an integer representing human years and returns an int in "dog years"
func Years(y int) int {
	return y * 7
}