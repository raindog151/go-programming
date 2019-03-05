package word

import (
	"fmt"
	"github.com/raindog151/go-programming/ninja_13/exercise_02/quote"
	"testing"
)

func TestCount(t *testing.T) {
	x := Count(quote.SunAlso)
	if x != 1349 {
		t.Error("Got", x, "Expected 1349")
	}
}

func TestUseCount(t *testing.T) {
	x := UseCount(quote.SunAlso)
	if x["magazine"] != 2 {
		t.Error("Got", x["magazine"], "Expected 3")
	}
}

func ExampleCount() {
	fmt.Println(Count(quote.SunAlso))
	// Output:
	// 1349
}

func ExampleUseCount() {
	x := UseCount(quote.SunAlso)
	fmt.Println(x["magazine"])
	// Output:
	// 2
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < 10; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < 10; i++ {
		UseCount(quote.SunAlso)
	}
}