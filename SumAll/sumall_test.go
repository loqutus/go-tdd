package sumall

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	want := []int{3, 7}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2}, []int{3, 4})
	}
}

func ExampleSumAll() {
	sum := SumAll([]int{1, 2}, []int{3, 4})
	fmt.Println(sum)
	// Output: [3 7]
}
