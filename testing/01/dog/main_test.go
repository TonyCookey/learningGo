package dog

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	x := Convert(7)
	if x != 49 {
		t.Error("Did not compute correct conversion")
	}

}

func ExampleConvert() {
	fmt.Println(Convert(7))
	//Output: 49
}

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Convert(7)
	}
}
