package adder

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("returns sum of two numbers", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("got %d want %d", sum, expected)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
