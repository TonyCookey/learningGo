package word

import (
	"fmt"
	"learningGo/testing/02/quote"
	"testing"
)

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}
func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func ExampleCount() {
	fmt.Println(Count(quote.SunAlso))
	//Output:
	//1349
}

func TestCount(t *testing.T) {
	val := Count(quote.SunAlso)
	if val != 1349 {
		t.Error("Expected", 1349, "Got", val)
	}
}

func TestUseCount(t *testing.T) {
	_, val := UseCount(quote.SunAlso)

	if val != 594 {
		t.Error("Expected", 594, "Got", val)
	}

}
