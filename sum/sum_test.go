package sum

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of fixed size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 2, 3, 4})
	}
}

func ExampleSum() {
	sum := Sum([]int{1, 2, 3, 4})
	fmt.Println(sum)
	// Output: 10
}
