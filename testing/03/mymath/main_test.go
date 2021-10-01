package mymath

import (
	"fmt"
	"testing"
)

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 4, 6, 8, 100})
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 4, 6, 8, 100}))
	//Output:
	// 6
}

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		answer float64
	}
	tests := []test{
		test{[]int{1, 4, 6, 8, 100}, float64(6)},
		test{[]int{0, 8, 10, 1000}, float64(9)},
	}

	for _, v := range tests {
		if CenteredAvg(v.data) != v.answer {
			t.Error("Expected", v.answer, "got", CenteredAvg(v.data))
		}
	}
}
func TestCenteredAvg2(t *testing.T) {
	x := CenteredAvg([]int{0, 8, 10, 1000})
	if x != 9 {
		t.Error("Expected", 9, "got", x)
	}
}
