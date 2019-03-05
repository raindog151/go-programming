package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	x := CenteredAvg([]int{1, 2, 3, 4, 5})
	if x != 3 {
		t.Error("Got", x, "Expected 3")
	}

	type set struct {
		numbers []int
		average float64
	}

	sets := []set {
		set{[]int{10, 20, 30, 40, 50}, 30},
		set{[]int{50, 60, 70, 80, 90}, 70},
		set{[]int{100, 200, 300, 400, 500}, 300},
	}

	for _, v := range sets {
		if CenteredAvg(v.numbers) != v.average {
			t.Error("Got", CenteredAvg(v.numbers), "Expected", v.average)
		}
	}

}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 3, 4, 5}))
	// Output:
	// 3
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < 10; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5})
	}
}
