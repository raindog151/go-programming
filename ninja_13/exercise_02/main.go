package main

import (
	"fmt"
	"github.com/raindog151/go-programming/ninja_13/exercise_02/quote"
	"github.com/raindog151/go-programming/ninja_13/exercise_02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
