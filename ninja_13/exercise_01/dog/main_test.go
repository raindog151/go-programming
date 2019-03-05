package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	x := Years(20)
	if x != 140 {
		t.Error("Got", x, "wanted 140")
	}
}

func TestYearsTwo(t *testing.T) {
	x := YearsTwo(20)
	if x != 140 {
		t.Error("Got", x, "wanted 140")
	}
}

func ExampleYears() {
	fmt.Println(Years(20))
	// Output:
	// 140
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(20))
	// Output:
	// 140
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < 20; i++ {
		Years(20)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < 20; i++ {
		YearsTwo(20)
	}
}